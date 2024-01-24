package templates

import (
	_ "embed"
	"text/template"
)

//go:embed cgo.go.template
var cgoTemplate []byte

// Cgo template (text/template)
func Cgo() (*template.Template, error) {

	return template.New("Cgo").Parse(string(cgoTemplate))
}
