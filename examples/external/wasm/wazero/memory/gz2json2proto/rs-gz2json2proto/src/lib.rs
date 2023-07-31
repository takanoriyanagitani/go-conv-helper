#![deny(unsafe_code)]

use bytes::BufMut;
use prost::Message;

static mut _INPUT: [u8; 65536] = [0; 65536];
static mut _OUTPUT: [u8; 65536] = [0; 65536];

#[derive(Debug)]
pub enum Error {
    UnableToSerialize,
    UnableToParse,
    Eof,
    Expected,
    InvalidNumber,
    InvalidType,
    InvalidUnicodeCodePoint,
    Key,
    Trailing,
    Custom,
    Unknown,
}

pub mod output {
    include!(concat!(env!("OUT_DIR"), "/output.v1.rs"));
}

pub fn slice2output(input: &[u8]) -> Result<output::Output, Error> {
    serde_json::from_slice(input).map_err(|_| Error::UnableToParse)
}

pub fn output2buf<M, B>(output: &M, buf: &mut B) -> Result<usize, Error>
where
    M: Message,
    B: BufMut,
{
    let sz: usize = output.encoded_len();
    output
        .encode(buf)
        .map(|_| sz)
        .map_err(|_| Error::UnableToSerialize)
}

fn serialize_output(o: &output::Output) -> Result<usize, Error> {
    #[allow(unsafe_code)]
    let mut out: &mut [u8] = unsafe { &mut _OUTPUT };
    output2buf(o, &mut out)
}

fn input() -> &'static [u8] {
    #[allow(unsafe_code)]
    unsafe {
        &_INPUT
    }
}

#[allow(unsafe_code)]
#[no_mangle]
pub extern "C" fn input_address() -> *mut u8 {
    unsafe { _INPUT.as_mut_ptr() }
}

#[allow(unsafe_code)]
#[no_mangle]
pub extern "C" fn output_address() -> *mut u8 {
    unsafe { _OUTPUT.as_mut_ptr() }
}

fn convert(input: &[u8]) -> i32 {
    slice2output(input)
        .and_then(|out| serialize_output(&out))
        .map(|sz: usize| sz as i32)
        .unwrap_or(-1)
}

#[allow(unsafe_code)]
#[no_mangle]
pub extern "C" fn converter(isz: i32) -> i32 {
    let i: &[u8] = input();
    let sz: usize = (isz & 0xffff) as usize;
    let limited: &[u8] = &i[..sz];
    convert(limited)
}

#[cfg(test)]
mod test {

    mod slice2output {
        #[test]
        fn test_empty() {
            let o: crate::output::Output = crate::slice2output(
                r#"{
                  "device": {
                      "id": {"id": "test"},
                      "name": {"name": "helo"}
                  },
                  "timestamp": {"unixtime_seconds": 0},
                  "current": {"current": 0.0},
                  "data": []
                }"#
                .as_bytes(),
            )
            .unwrap();
        }
    }
}
