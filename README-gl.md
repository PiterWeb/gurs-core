# gurs-core

## 🤯 ¿Qué é gurs-core?

Librería que fai unha análise de código de rust e crea unha interfaz para usar nos teus proxectos de golang
(está pensada para ser usada por unha cli que funcione en windows/linux/osx):

- Usando https://wazero.io/ [Necesitase compilar a Webassembly]

- Usando DLL / SO / LYB (Librerías dependentes do sistema, podense cargar sen paquete) [Necesitase usar comandos de rustc]
(Crear Rust DLL)
https://samrambles.com/guides/window-hacking-with-rust/creating-a-dll-with-rust/index.html#hellodll
(Cargar DLL go // Posíbel Opción)
https://github.com/ebitengine/purego

- Usando CGo:
(Compilar Rust a unha librería de C e xerar os .h) https://github.com/mozilla/cbindgen
[Ejemplo de uso] https://github.com/getsentry/milksnake

Utilizar text/template para ter unhas plantillas co código de golang que se embeberan na build do paquete
e que en tempo de execución se sustituiran os valores necesarios para que corra segúndo as opcións escollidas posteriormente na cli

Máis Documentación:

- Rust with C: https://docs.rust-embedded.org/book/interoperability/rust-with-c.html#no_mangle

