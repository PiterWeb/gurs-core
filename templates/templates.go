package templates

import (
	_ "embed"
)

//go:embed cgo.go.txt
var cgoTemplate []byte

//go:embed dll.go.txt
var dllTemplate []byte

//go:embed wazero.go.txt
var wazeroTemplate []byte

func Cgo() string {
	return string(cgoTemplate)
}

func Dll() string {
	return string(dllTemplate)
}

func Wazero() string {
	return string(wazeroTemplate)
}
