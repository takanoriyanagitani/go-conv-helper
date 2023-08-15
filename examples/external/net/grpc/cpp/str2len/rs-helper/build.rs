use std::io;

use tonic_build::Builder;

fn main() -> Result<(), io::Error> {
    let b: Builder = tonic_build::configure().build_server(true);
    b.compile(&["proto/str2len/v1/sl.proto"], &["proto/str2len/v1"])?;
    Ok(())
}
