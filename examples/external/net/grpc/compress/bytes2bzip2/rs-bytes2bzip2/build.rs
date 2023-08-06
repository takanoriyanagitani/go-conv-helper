use std::io;

use tonic_build::Builder;

fn main() -> Result<(), io::Error> {
    let b: Builder = tonic_build::configure().build_server(true);
    b.compile(&["b2b.proto"], &["proto/bytes2bz/v1"])?;
    Ok(())
}
