package templates

import (
	_ "embed"
	"fmt"
	"strings"
	"text/template"

	"github.com/PiterWeb/gurs-core/parser"
	"github.com/PiterWeb/gurs-core/transpile"
)

type CgoTemplate struct {
	Functions   []parser.Gofunc
	Package     string
	GursVersion string
}

//go:embed cgo.go.tpl
var cgoTemplate []byte

// Cgo template (text/template)
func Cgo() (*template.Template, error) {

	funcsMap := template.FuncMap{
		"ReplaceFnBody":    ReplaceFnBody,
		"CreateVariables":  CreateVariables,
		"GetVariableNames": GetVariableNames,
	}

	return template.New("Cgo").Funcs(funcsMap).Parse(string(cgoTemplate))
}

func ReplaceFnBody(fn string, body ...string) string {

	return strings.Replace(fn, "{{.}}", strings.Join(body, "\n"), 1)

}

func GetVariableNames(fn parser.Gofunc) string {

	variables := ""

	for i, p := range fn.Parameters {
		// If it's the last parameter, don't add a comma
		if i == len(fn.Parameters)-1 {
			variables += fmt.Sprintf("c_%s", p.GetName())
			break
		}
		variables += fmt.Sprintf("c_%s, ", p.GetName())
	}

	return variables

}

func CreateVariables(fn parser.Gofunc) string {

	variables := "\n"

	for _, p := range fn.Parameters {
		cgo_type := transpile.GoTypeToCGoType(p.GetType())
		arg_name := p.GetName()
		variables += fmt.Sprintf("c_%s := %s(%s)\n", arg_name, cgo_type, arg_name)
	}

	return variables
}
