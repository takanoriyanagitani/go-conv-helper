use std::path::{Path, PathBuf};

#[cfg(not(unix))]
use std::net::SocketAddr;

use tonic::transport::{server::Router, Server};

#[cfg(unix)]
use tokio::net::UnixListener;

#[cfg(unix)]
use tokio_stream::wrappers::UnixListenerStream;

use rs_helper::proxy::event::send_event_service_new;
use rs_helper::proxy::get::{cmd_src_new, get_cmd_svc_new};
use rs_helper::s2l::get_cmd_service_server::GetCmdServiceServer;
use rs_helper::s2l::proxy::get_cmd_response::Command;
use rs_helper::s2l::proxy::{ReplyId, StringLengthCommand, Uuid};
use rs_helper::s2l::send_event_service_server::SendEventServiceServer;
use rs_helper::s2l::string_length_service_server::StringLengthServiceServer;
use rs_helper::s2l::StringLengthRequest;

#[cfg(not(unix))]
async fn sub() -> Result<(), String> {
    let addr_str: String =
        std::env::var("ENV_ADDR").unwrap_or_else(|_| String::from("127.0.0.1:50051"));
    let dummy_commands: Vec<Command> = vec![
        Command::StrlenCmd(StringLengthCommand {
            strlen_req_id: None,
            reply: Some(ReplyId {
                id: Some(Uuid {
                    hi: 0xcafef00ddeadbeaf,
                    lo: 0xface864299792458,
                }),
            }),
            request: Some(StringLengthRequest {
                source: String::from("count me"),
            }),
        }),
        Command::StrlenCmd(StringLengthCommand {
            strlen_req_id: None,
            reply: None,
            request: None,
        }),
    ];
    let mut dcmd_iter = dummy_commands.into_iter();
    let dummy_cmd_src = cmd_src_new(move |_| Ok(dcmd_iter.next()));
    let cmd_svc = get_cmd_svc_new(dummy_cmd_src);
    let cmd_srv: GetCmdServiceServer<_> = GetCmdServiceServer::new(cmd_svc);

    let dummy_send_svc = send_event_service_new();
    let dummy_send_svr: SendEventServiceServer<_> = SendEventServiceServer::new(dummy_send_svc);

    let svc = rs_helper::str_len_svc_new();
    let srv: StringLengthServiceServer<_> = StringLengthServiceServer::new(svc);

    let addr: SocketAddr =
        str::parse(addr_str.as_str()).map_err(|e| format!("Invalid addr: {e}"))?;

    let mut gsv: Server = Server::builder();
    let router: Router<_> = gsv
        .add_service(srv)
        .add_service(cmd_srv)
        .add_service(dummy_send_svr);
    router
        .serve(addr)
        .await
        .map_err(|e| format!("Unable to start gRPC server: {e}"))?;
    Ok(())
}

#[cfg(unix)]
async fn sub() -> Result<(), String> {
    const PATH2SOCKDIR: &str = "./sock.d";
    std::fs::create_dir_all(Path::new(PATH2SOCKDIR))
        .map_err(|e| format!("Unable to create a dir: {e}"))?;
    let path2sock: PathBuf = Path::new(PATH2SOCKDIR).join("test.sock");
    std::fs::remove_file(&path2sock).ok();

    let uds: UnixListener = UnixListener::bind(path2sock).unwrap();
    let uds_stream: UnixListenerStream = UnixListenerStream::new(uds);

    let mut gsv: Server = Server::builder();

    let dummy_commands: Vec<Command> = vec![
        Command::StrlenCmd(StringLengthCommand {
            strlen_req_id: None,
            reply: Some(ReplyId {
                id: Some(Uuid {
                    hi: 0xcafef00ddeadbeaf,
                    lo: 0xface864299792458,
                }),
            }),
            request: Some(StringLengthRequest {
                source: String::from("count me"),
            }),
        }),
        Command::StrlenCmd(StringLengthCommand {
            strlen_req_id: None,
            reply: None,
            request: None,
        }),
    ];

    let mut dcmd_iter = dummy_commands.into_iter();
    let dummy_cmd_src = cmd_src_new(move |_| Ok(dcmd_iter.next()));
    let cmd_svc = get_cmd_svc_new(dummy_cmd_src);
    let cmd_srv: GetCmdServiceServer<_> = GetCmdServiceServer::new(cmd_svc);

    let dummy_send_svc = send_event_service_new();
    let dummy_send_svr: SendEventServiceServer<_> = SendEventServiceServer::new(dummy_send_svc);
    let svc = rs_helper::str_len_svc_new();
    let srv: StringLengthServiceServer<_> = StringLengthServiceServer::new(svc);

    let router: Router<_> = gsv
        .add_service(srv)
        .add_service(cmd_srv)
        .add_service(dummy_send_svr);

    router
        .serve_with_incoming(uds_stream)
        .await
        .map_err(|e| format!("Unable to listen: {e}"))?;
    Ok(())
}

#[tokio::main]
async fn main() -> Result<(), String> {
    sub().await
}
