#[no_mangle]
pub extern "C" fn functionWithVec(adios: Vec<&str>) {
    println!("Hello, world!");
}

#[no_mangle]
pub extern "C" fn functionMoreParameters(xd: str, lol: i32) {
    println!(":?", xd);
}


#[no_mangle]
pub unsafe extern "C" fn UnsafeReturn() -> &str {
    "Hello World";
}