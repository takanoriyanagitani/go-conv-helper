#![deny(unsafe_code)]

#[allow(unsafe_code)]
#[no_mangle]
pub extern "C" fn sinf32(input: f32) -> f32 {
    input.sin()
}

#[allow(unsafe_code)]
#[no_mangle]
pub extern "C" fn sinf64(input: f64) -> f64 {
    input.sin()
}
