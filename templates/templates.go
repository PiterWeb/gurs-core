package templates

import (
	_ "embed"
	"text/template"
)

//go:embed cgo.go.template
var cgoTemplate []byte

//go:embed wazero.go.template
var wazeroTemplate []byte

//go:embed wazero_runtime.go.template
var wazeroRuntime []byte

// Cgo template (text/template)
func Cgo() (*template.Template, error) {

	return template.New("Cgo").Parse(string(cgoTemplate))
}

// Wazero template (text/template)
func Wazero() (*template.Template, error) {
	return template.New("Wazero").Parse(string(wazeroTemplate))
}

// Wazero runtime file (no need to edit so is a string)
func WazeroRuntime() string {
	return string(wazeroRuntime)
}
