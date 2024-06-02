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
		rawParameters = append(rawParameters, fmt.Sprintf("%s %s", p.Name, p.ValueType))
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

func (fn Gofunc) ToTemplate() (*template.Template, error) {

	templ := template.New("func-" + fn.Name)

	textTemplate := fn.ToString()

	templ, err := templ.Parse(textTemplate)

	if err != nil {
		return nil, err
	}

	return templ, nil

}

func (fn Gofunc) ToString() string {

	// Convert the parameters to a string
	parameters := strings.Join(baseFunction(fn).GetRawParameters(), ", ")

	// Replace the colon with a space
	parameters = strings.ReplaceAll(parameters, ":", " ")

	textTemplate := fmt.Sprintf("func %s(%s) %s {{{.}}}\n", fn.Name, parameters, fn.ReturnType)

	return textTemplate
}

// func (fn Gofunc) CGo() string {

// 	params := [][]string{}

// 	for _, fullParam := range baseFunction(fn).GetRawParameters() {

// 		param := strings.Split(fullParam, " ")

// 		params = append(params, param)

// 	}

// 	transpiled := transpile.Cgo(fn.Name, fn.ReturnType, params)

// 	return transpiled

// }

func ConvertRsFnSliceToGo(fns []Rustfn) (goFuncs []Gofunc) {

	for _, r := range fns {

		goFuncs = append(goFuncs, r.ConvertToGo())

	}

	return

}
