#![deny(unsafe_code)]

use std::io::{Read, Write};

use flate2::read::GzDecoder;

static mut _INPUT: [u8; 65536] = [0; 65536];
static mut _OUTPUT: [u8; 65536] = [0; 65536];

fn input() -> &'static [u8] {
    #[allow(unsafe_code)]
    unsafe {
        &_INPUT
    }
}

fn output() -> &'static mut [u8] {
    #[allow(unsafe_code)]
    unsafe {
        &mut _OUTPUT
    }
}

#[allow(unsafe_code)]
#[no_mangle]
pub extern "C" fn input_address() -> *mut u8 {
    #[allow(unsafe_code)]
    unsafe {
        _INPUT.as_mut_ptr()
    }
}

#[allow(unsafe_code)]
#[no_mangle]
pub extern "C" fn output_address() -> *mut u8 {
    #[allow(unsafe_code)]
    unsafe {
        _OUTPUT.as_mut_ptr()
    }
}

fn dummy_count_len<R>(r: R) -> usize
where
    R: Read,
{
    let bytes = r.bytes();
    bytes.count()
}

fn dummy_count_slice_via_read(i: &[u8]) -> usize {
    dummy_count_len(i)
}

#[allow(unsafe_code)]
#[no_mangle]
pub extern "C" fn dummy_count() -> usize {
    let i: &[u8] = input();
    dummy_count_slice_via_read(i)
}

fn reader2writer<R, W>(mut reader: R, mut writer: W) -> Result<u64, String>
where
    R: Read,
    W: Write,
{
    let rdr = reader.by_ref();
    let wtr = writer.by_ref();
    std::io::copy(rdr, wtr).map_err(|e| format!("Unable to copy: {e}"))
}

fn reader2output<R>(reader: R, output: &mut [u8]) -> Result<u64, String>
where
    R: Read,
{
    reader2writer(reader, output)
}

fn gz2output<R>(original: R, output: &mut [u8]) -> Result<u64, String>
where
    R: Read,
{
    let reader: GzDecoder<R> = GzDecoder::new(original);
    reader2output(reader, output)
}

fn input2output(input: &[u8], output: &mut [u8]) -> Result<u64, String> {
    gz2output(input, output)
}

#[allow(unsafe_code)]
#[no_mangle]
pub extern "C" fn converter(isz: i32) -> i32 {
    let i: &[u8] = input();
    let o: &mut [u8] = output();
    let sz: usize = isz as usize;
    let limited: &[u8] = &i[..sz];
    input2output(limited, o)
        .map(|sz: u64| sz as i32)
        .unwrap_or(-1)
}
