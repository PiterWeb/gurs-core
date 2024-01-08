#[no_mangle]
pub extern "C" fn HelloWorld() {
    println!("Hello, world!");
}

#[no_mangle]
pub extern "C" fn Caca(xd str) {
    println!(":?", xd);
}

#[no_mangle]
pub unsafe extern "C" fn UnsafeHelloWorld() -> &str {
    "Hello World";
}