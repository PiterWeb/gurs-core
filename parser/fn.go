package parser

import (
	"fmt"
	"strings"
	"text/template"

	"github.com/PiterWeb/gurs-core/transpile"
)

type baseFunction struct {
	FileName   string      `json:"fileName"`
	Name       string      `json:"name"`
	Parameters []parameter `json:"parameters"`
	ReturnType string      `json:"returnType"`
}

type parameter struct {
	Name      string `json:"name"`
	ValueType string `json:"valueType"`
}

type Rustfn baseFunction

func (f baseFunction) GetFileName() string {
	return f.FileName
}

func (f baseFunction) GetName() string {
	return f.Name
}

func (f baseFunction) GetParameters() []parameter {
	return f.Parameters
}

func (f baseFunction) GetRawParameters() []string {

	rawParameters := []string{}

	for _, p := range f.Parameters {
		rawParameters = append(rawParameters, fmt.Sprintf("%s: %s", p.Name, p.ValueType))
	}

	return rawParameters
}

func (f baseFunction) GetReturnType() string {
	return f.ReturnType
}

func (p parameter) GetName() string {
	return p.Name
}

func (p parameter) GetType() string {
	return p.ValueType
}

type Gofunc baseFunction

func (fn Rustfn) ConvertToGo() Gofunc {

	convertedParameters := []parameter{}

	for _, p := range baseFunction(fn).GetParameters() {

		convertedParameters = append(convertedParameters, parameter{
			Name:      p.Name,
			ValueType: transpile.TranspileType(p.ValueType),
		})

	}

	// Capitalize the word to make the function public
	publicFunctionName := strings.ToUpper(string(fn.Name[0])) + fn.Name[1:]

	f := Gofunc{
		FileName:   fn.FileName,
		Name:       publicFunctionName,
		Parameters: convertedParameters,
		ReturnType: transpile.TranspileType(fn.ReturnType),
	}

	return f

}

func (fn Gofunc) ToStringTemplate() (*template.Template, error) {

	templ := template.New("func-" + fn.Name)

	parameters := strings.Join(baseFunction(fn).GetRawParameters(), ", ")

	textTemplate := fmt.Sprintf("func %s(%s) %s {\n {{.}} \n}\n", fn.FileName, parameters, fn.ReturnType)

	templ, err := templ.Parse(textTemplate)

	if err != nil {
		return nil, err
	}

	return templ, nil

}

func ConvertRsFnSliceToGo(fns []Rustfn) (goFuncs []Gofunc) {

	for _, r := range fns {

		goFuncs = append(goFuncs, r.ConvertToGo())

	}

	return

}
