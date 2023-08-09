use std::net::SocketAddr;

use tonic::transport::{server::Router, Server};

use rs_bytes2bzip2::by2bz::compress_service_server::CompressServiceServer;
use rs_bytes2bzip2::compress::srv::compress_service_new;

#[tokio::main]
async fn main() -> Result<(), String> {
    let csvc: CompressServiceServer<_> = CompressServiceServer::new(compress_service_new(1024));
    let server: Server = Server::builder();
    let router: Router<_> = server
        .accept_http1(true)
        .add_service(tonic_web::enable(csvc));
    let addr: SocketAddr =
        str::parse("127.0.0.1:50051").map_err(|e| format!("Invalid listen addr: {e}"))?;
    router
        .serve(addr)
        .await
        .map_err(|e| format!("Unable to listen: {e}"))?;
    Ok(())
}
