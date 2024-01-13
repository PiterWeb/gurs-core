# Gurs Core
[![Go Reference](https://pkg.go.dev/badge/github.com/PiterWeb/gurs-core.svg)](https://pkg.go.dev/github.com/PiterWeb/gurs-core)

> [!WARNING]  
> Still on development ... 🚧

### 🏴 Languages

- [Español](/README-es.md)
- [Русский](/README-ru.md)
- [Galego](/README-gl.md)

## Prerequisites for use:

- Go
- Cargo
- CGo (only in project)

## 🤯 What is gurs-core ?

Library that parses rust code & create an interface to use on your golang projects
(intended to be used by a cli running on windows/linux/osx):

- Using https://wazero.io/ [Needs to compile to Webassembly].

- Using DLL / SO / LYB (System Dependent Libraries, which can be loaded without a package) [Requires using rustc commands].
(Create Rust DLL)
https://samrambles.com/guides/window-hacking-with-rust/creating-a-dll-with-rust/index.html#hellodll
(Load DLL go // Possible Option)
https://github.com/ebitengine/purego

- Using CGo:
(Compile Rust to C library and generate .h) https://github.com/mozilla/cbindgen
[Example of use] https://github.com/getsentry/milksnake

Use text/template to have a template with golang code that will be embedded in the build of the package
and that at runtime will substitute the necessary values to make it run according to the options chosen later in the cli

More Documentation:

- Rust with C: https://docs.rust-embedded.org/book/interoperability/rust-with-c.html#no_mangle
- Rust struct methods with C: https://stackoverflow.com/questions/54156498/how-to-call-a-rust-structs-method-from-c-using-ffi
- How to use C libraries in Go with CGo: https://dev.to/metal3d/understand-how-to-use-c-libraries-in-go-with-cgo-3dbn
- Specify Rust build target: https://doc.rust-lang.org/cargo/reference/config.html#buildtarget

## OS Support

| Windows 	| Linux 	| MacOS 	| FreeBSD 	|
|---------	|-------	|-------	|---------	|
| ✔       	| ✔     	| ✔     	| ❓       	|
