use std::io;

fn main() -> Result<(), io::Error> {
    let mut cfg = prost_build::Config::new();
    cfg.btree_map(["."]);
    cfg.type_attribute(".", "#[derive(serde::Deserialize)]");
    cfg.compile_protos(&["./src/out.proto"], &["./src/"])?;
    Ok(())
}
