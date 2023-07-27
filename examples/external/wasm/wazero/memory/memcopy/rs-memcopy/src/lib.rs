#![deny(unsafe_code)]

static mut _INPUT: [u8; 65536] = [0; 65536];
static mut _OUTPUT: [u8; 65536] = [0; 65536];

#[allow(unsafe_code)]
#[no_mangle]
pub extern "C" fn input_address() -> *mut u8 {
    unsafe { _INPUT.as_mut_ptr() }
}

#[allow(unsafe_code)]
#[no_mangle]
pub extern "C" fn output_address() -> *mut u8 {
    #[allow(unsafe_code)]
    unsafe {
        _OUTPUT.as_mut_ptr()
    }
}

fn convert_nocheck(i: &[u8], out: &mut [u8]) -> i32 {
    out.copy_from_slice(i);
    out.len() as i32
}

fn _input() -> &'static [u8] {
    #[allow(unsafe_code)]
    unsafe {
        &_INPUT
    }
}
fn _output() -> &'static mut [u8] {
    #[allow(unsafe_code)]
    unsafe {
        &mut _OUTPUT
    }
}

#[allow(unsafe_code)]
#[no_mangle]
pub extern "C" fn converter(isz: i32) -> i32 {
    let i: &[u8] = _input();
    let o: &mut [u8] = _output();
    let sz: usize = (isz & 0xffff) as usize;
    convert_nocheck(&i[..sz], &mut o[..sz])
}
