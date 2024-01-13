#[no_mangle]
pub extern "C" fn ParamVec(adios: Vec<&str>) {
    println!("Hello, world!");
}

#[no_mangle]
pub extern "C" fn TwoParams(xd: str, lol: i32) {
    println!(":?", xd);
}


#[no_mangle]
pub unsafe extern "C" fn UnsafeReturn() -> &str {
    "Hello World";
}