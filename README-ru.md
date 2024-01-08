# gurs-core
[![Go Reference](https://pkg.go.dev/badge/github.com/PiterWeb/gurs-core.svg)](https://pkg.go.dev/github.com/PiterWeb/gurs-core)

## 🤯 Что такое gurs-core ?

Библиотека, которая анализирует код rust и создает интерфейс для использования в ваших проектах golang
(предназначена для использования cli, работающим в Windows/linux/osx):

- С помощью https://wazero.io/ [Необходимо скомпилировать в Webassembly].

- Использование DLL / SO / LYB (системно-зависимые библиотеки, которые могут быть загружены без пакета) [Требуется использование команд rustc].
(Создать Rust DLL)
https://samrambles.com/guides/window-hacking-with-rust/creating-a-dll-with-rust/index.html#hellodll
(Загрузить DLL go // Возможный вариант)
https://github.com/ebitengine/purego

- Использование CGo:
(Скомпилируйте Rust в библиотеку C и сгенерируйте .h) https://github.com/mozilla/cbindgen
[Пример использования] https://github.com/getsentry/milksnake

Используйте text/template, чтобы создать шаблон с кодом golang, который будет встроен в сборку пакета
и который во время выполнения заменит необходимые значения, чтобы заставить его работать в соответствии с параметрами, выбранными позже в cli

Дополнительная документация:

- Rust with C: https://docs.rust-embedded.org/book/interoperability/rust-with-c.html#no_mangle
- Rust struct methods with C: https://stackoverflow.com/questions/54156498/how-to-call-a-rust-structs-method-from-c-using-ffi

