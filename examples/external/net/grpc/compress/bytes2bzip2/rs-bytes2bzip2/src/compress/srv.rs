use std::io;
use std::io::{Read, Write};

use tonic::{Request, Response, Status};

use bzip2::{write::BzEncoder, Compression};

use crate::by2bz::compress_service_server::CompressService;
use crate::by2bz::compression::Level;
use crate::by2bz::{CompressRequest, CompressResponse};

pub struct CompressSvc {
    capacity: usize,
}

impl CompressSvc {
    pub fn convert_level(l: Level) -> Result<Compression, String> {
        match l {
            Level::NumFast => Ok(Compression::fast()),
            Level::NumBest => Ok(Compression::best()),
            Level::Num0 => Ok(Compression::new(0)),
            Level::Num1 => Ok(Compression::new(1)),
            Level::Num2 => Ok(Compression::new(2)),
            Level::Num3 => Ok(Compression::new(3)),
            Level::Num4 => Ok(Compression::new(4)),
            Level::Num5 => Ok(Compression::new(5)),
            Level::Num6 => Ok(Compression::new(6)),
            Level::Num7 => Ok(Compression::new(7)),
            Level::Num8 => Ok(Compression::new(8)),
            Level::Num9 => Ok(Compression::new(9)),
            Level::NumUnspecified => Err(format!("Unknown level: {l:?}")),
        }
    }

    pub fn compress<R, W>(level: Compression, mut i: R, o: W) -> Result<(), String>
    where
        R: Read,
        W: Write,
    {
        let mut encoder: BzEncoder<_> = BzEncoder::new(o, level);
        let rdr: &mut R = &mut i;
        io::copy(rdr, &mut encoder)
            .map(|_| ())
            .map_err(|e| format!("Unable to compress: {e}"))
    }

    pub fn compress2buf<R>(level: Compression, i: R, buf: &mut Vec<u8>) -> Result<(), String>
    where
        R: Read,
    {
        buf.clear();
        Self::compress(level, i, buf)
    }

    pub fn slice2buf(level: Compression, i: &[u8], buf: &mut Vec<u8>) -> Result<(), String> {
        Self::compress2buf(level, i, buf)
    }
}

#[tonic::async_trait]
impl CompressService for CompressSvc {
    async fn compress(
        &self,
        req: Request<CompressRequest>,
    ) -> Result<Response<CompressResponse>, Status> {
        let q: &CompressRequest = req.get_ref();
        let l: Level = q.level();
        let c: Compression = Self::convert_level(l)
            .map_err(|e| Status::invalid_argument(format!("invalid argument: {e}")))?;
        let mut buf: Vec<u8> = Vec::with_capacity(self.capacity);
        let input: &[u8] = &q.input;
        Self::slice2buf(c, input, &mut buf)
            .map_err(|e| Status::internal(format!("Unable to compress: {e}")))?;
        let reply: CompressResponse = CompressResponse { compressed: buf };
        Ok(Response::new(reply))
    }
}

pub fn compress_service_new(capacity: usize) -> impl CompressService {
    CompressSvc { capacity }
}

pub fn compress_service_new_empty() -> impl CompressService {
    compress_service_new(0)
}
