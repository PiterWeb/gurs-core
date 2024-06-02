#[no_mangle]
pub extern "C" fn paramVec(adios: Vec<&str>) {
    println!("Hello, world!");
}

#[no_mangle]
pub extern "C" fn otherParamVec(adios: Vec<i32>) {
    println!("lel");
}

#[no_mangle]
pub extern "C" fn twoParams(xd: str, _lol: i32) {
    println!(":?", xd);
}

#[no_mangle]
pub unsafe extern "C" fn unsafeReturn() -> &str {
    "Hello World";
}