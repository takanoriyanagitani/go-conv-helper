use tonic::{Request, Response, Status};

pub mod s2l {
    tonic::include_proto!("str2len.v1");
}

pub mod proxy;

use crate::s2l::string_length_service_server::StringLengthService;
use crate::s2l::{StringLengthRequest, StringLengthResponse};

#[derive(Default)]
pub struct StrLenSvc {}

#[tonic::async_trait]
impl StringLengthService for StrLenSvc {
    async fn string_length(
        &self,
        req: Request<StringLengthRequest>,
    ) -> Result<Response<StringLengthResponse>, Status> {
        let r: &StringLengthRequest = req.get_ref();
        let source: &str = r.source.as_str();
        let length: u64 = source.len() as u64;
        let reply: StringLengthResponse = StringLengthResponse { length };
        Ok(Response::new(reply))
    }
}

pub fn str_len_svc_new() -> impl StringLengthService {
    StrLenSvc::default()
}
