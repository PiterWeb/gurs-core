package templates

import (
	_ "embed"
	"text/template"
)

//go:embed wazero.go.tpl
var wazeroTemplate []byte

//go:embed wazero_runtime.go.tpl
var wazeroRuntime []byte

// Wazero template (text/template)
func Wazero() (*template.Template, error) {
	return template.New("Wazero").Parse(string(wazeroTemplate))
}

// Wazero runtime file (no need to edit so is a string)
func WazeroRuntime() string {
	return string(wazeroRuntime)
}
