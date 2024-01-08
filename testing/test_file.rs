#[no_mangle]
pub extern "C" fn HelloWorld(adios: Vec<&str>) {
    println!("Hello, world!");
}

#[no_mangle]
pub extern "C" fn Caca(xd: str, lol: i32) {
    println!(":?", xd);
}


#[no_mangle]
pub unsafe extern "C" fn UnsafeHelloWorld() -> &str {
    "Hello World";
}