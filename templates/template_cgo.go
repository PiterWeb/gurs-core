package templates

import (
	_ "embed"
	"strings"
	"text/template"

	"github.com/PiterWeb/gurs-core/parser"
)

type CgoTemplate struct {
	Functions   []parser.GoFn
	Package     string
	GursVersion string
}

//go:embed cgo.go.tpl
var cgoTemplate []byte

// Cgo template (text/template)
func Cgo() (*template.Template, error) {

	funcsMap := template.FuncMap{
		"ReplaceFnBody": ReplaceFnBody,
	}

	return template.New("Cgo").Funcs(funcsMap).Parse(string(cgoTemplate))
}

func ReplaceFnBody(fn string, body ...string) string {

	return strings.Replace(fn, "{{.}}", strings.Join(body, "\n"), 1)

}
