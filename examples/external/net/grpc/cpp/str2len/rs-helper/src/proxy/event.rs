use std::time::SystemTime;

use prost_types::Timestamp;
use tonic::{Request, Response, Status};

use crate::s2l::proxy::{SendEventRequest, SendEventResponse};
use crate::s2l::send_event_service_server::SendEventService;

#[derive(Default)]
pub struct SendEvtSvc {}

#[tonic::async_trait]
impl SendEventService for SendEvtSvc {
    async fn send_event(
        &self,
        _request: Request<SendEventRequest>,
    ) -> Result<Response<SendEventResponse>, Status> {
        let now: SystemTime = SystemTime::now();
        let ts: Timestamp = Timestamp::try_from(now)
            .map_err(|e| Status::internal(format!("Unexpected error: {e}")))?;
        let reply: SendEventResponse = SendEventResponse { sent: Some(ts) };
        Ok(Response::new(reply))
    }
}

pub fn send_event_service_new() -> impl SendEventService {
    SendEvtSvc::default()
}
