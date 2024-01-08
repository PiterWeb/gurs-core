# gurs-core
[![Go Reference](https://pkg.go.dev/badge/github.com/PiterWeb/gurs-core.svg)](https://pkg.go.dev/github.com/PiterWeb/gurs-core)

## 🤯 ¿Qué es gurs-core?

Librería que analiza código de rust y crea una interfaz para usar en tus proyectos de golang
(está pensada para ser usada por una cli que funcione en windows/linux/osx):

- Usando https://wazero.io/ [Se necesita compilar a Webassembly]

- Usando DLL / SO / LYB (Librerías dependientes de sistema, que se pueden cargar sin paquete) [Se necesita usar comandos de rustc]
(Crear Rust DLL)
https://samrambles.com/guides/window-hacking-with-rust/creating-a-dll-with-rust/index.html#hellodll
(Cargar DLL go // Posible Opción)
https://github.com/ebitengine/purego

- Usando CGo:
(Compilar Rust a librería de C y generar .h) https://github.com/mozilla/cbindgen
[Ejemplo de uso] https://github.com/getsentry/milksnake

Utilizar text/template para tener unas template con código de golang que se embeberan en la build del paquete
y que en tiempo de ejecución se sustituiran los valores necesarios para que corra según las opciones escogidas posteriormente en el cli

Más Documentación:

- Rust with C: https://docs.rust-embedded.org/book/interoperability/rust-with-c.html#no_mangle
- Rust struct methods with C: https://stackoverflow.com/questions/54156498/how-to-call-a-rust-structs-method-from-c-using-ffi

