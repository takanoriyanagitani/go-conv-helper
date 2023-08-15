use core::time::Duration;
use std::sync::Mutex;

use tokio::time::Interval;
use tonic::{Request, Response, Status};

use crate::s2l::get_cmd_service_server::GetCmdService;
use crate::s2l::proxy::get_cmd_response::Command;
use crate::s2l::proxy::{GetCmdRequest, GetCmdResponse, Retry, Target, Uuid};

const WAIT_DEFAULT: prost_types::Duration = prost_types::Duration {
    seconds: 0,
    nanos: 1_000_000,
};

const RETRY_DEFAULT: Retry = Retry {
    max_retry_count: 10,
};

pub trait CommandSource {
    fn get_command(&self, id: &Uuid) -> Result<Option<Command>, String>;
}

struct CmdSrc<F> {
    internal: Mutex<F>,
}
impl<F> CommandSource for CmdSrc<F>
where
    F: FnMut(&Uuid) -> Result<Option<Command>, String>,
{
    fn get_command(&self, id: &Uuid) -> Result<Option<Command>, String> {
        match self.internal.lock() {
            Err(e) => Err(format!("Unable to lock: {e}")),
            Ok(mut guard) => {
                let mf: &mut F = &mut guard;
                mf(id)
            }
        }
    }
}

pub fn cmd_src_new<F>(src_fn: F) -> impl CommandSource
where
    F: FnMut(&Uuid) -> Result<Option<Command>, String>,
{
    CmdSrc {
        internal: Mutex::new(src_fn),
    }
}

pub struct GetCmdSvc<S> {
    cmd_src: S,
}

pub fn get_cmd_svc_new<S>(cmd_src: S) -> impl GetCmdService
where
    S: CommandSource + Sync + Send + 'static,
{
    GetCmdSvc { cmd_src }
}

impl<S> GetCmdSvc<S>
where
    S: CommandSource,
{
    pub fn get_command(&self, id: &Uuid) -> Result<Option<Command>, Status> {
        self.cmd_src
            .get_command(id)
            .map_err(|e| Status::internal(format!("Unable to get a command: {e}")))
    }
}

#[tonic::async_trait]
impl<S> GetCmdService for GetCmdSvc<S>
where
    S: CommandSource + Send + Sync + 'static,
{
    async fn get_cmd(
        &self,
        req: Request<GetCmdRequest>,
    ) -> Result<Response<GetCmdResponse>, Status> {
        let r: &GetCmdRequest = req.get_ref();
        let wait: prost_types::Duration = r.wait.clone().unwrap_or(WAIT_DEFAULT);
        let wt: Duration = wait
            .try_into()
            .map_err(|e| Status::invalid_argument(format!("invalid wait: {e}")))?;
        let mut invl: Interval = tokio::time::interval(wt);
        let retry: Retry = r.retry.clone().unwrap_or(RETRY_DEFAULT);
        let target: &Target = r
            .target
            .as_ref()
            .ok_or_else(|| Status::invalid_argument("target missing"))?;
        let uuid: Uuid = target
            .id
            .clone()
            .ok_or_else(|| Status::invalid_argument("target uuid missing"))?;
        for _ in 0..retry.max_retry_count {
            invl.tick().await; // 1st tick has 0 latency
            match self.get_command(&uuid) {
                Err(e) => {
                    return Err(e);
                }
                Ok(Some(cmd)) => {
                    let reply: GetCmdResponse = GetCmdResponse { command: Some(cmd) };
                    return Ok(Response::new(reply));
                }
                Ok(None) => {}
            }
        }
        Ok(Response::new(GetCmdResponse {
            command: Some(Command::Nop(())),
        }))
    }
}
