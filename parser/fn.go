package parser

import (
	"fmt"
	"strings"
	"text/template"

	"github.com/PiterWeb/gurs-core/transpile"
)

type baseFunction struct {
	fileName   string
	name       string
	parameters []parameter
	returnType string
}

type Rustfn baseFunction

type parameter struct {
	name      string
	valueType string
}

func (f baseFunction) GetFileName() string {
	return f.fileName
}

func (f baseFunction) GetName() string {
	return f.name
}

func (f baseFunction) GetParameters() []parameter {
	return f.parameters
}

func (f baseFunction) GetRawParameters() []string {

	rawParameters := []string{}

	for _, p := range f.parameters {
		rawParameters = append(rawParameters, fmt.Sprintf("%s: %s", p.name, p.valueType))
	}

	return rawParameters
}

func (f baseFunction) GetReturnType() string {
	return f.returnType
}

func (p parameter) GetName() string {
	return p.name
}

func (p parameter) GetType() string {
	return p.valueType
}

type Gofunc baseFunction

func (fn Rustfn) ConvertToGo() Gofunc {

	convertedParameters := []parameter{}

	for _, p := range baseFunction(fn).GetParameters() {

		convertedParameters = append(convertedParameters, parameter{
			name:      p.name,
			valueType: transpile.TranspileType(p.valueType),
		})

	}

	f := Gofunc{
		fileName:   fn.fileName,
		name:       fn.name,
		parameters: convertedParameters,
		returnType: transpile.TranspileType(fn.returnType),
	}

	return f

}

func (fn Gofunc) ToStringTemplate() (*template.Template, error) {

	templ := template.New("func-" + fn.name)

	parameters := strings.Join(baseFunction(fn).GetRawParameters(), ", ")

	// Capitalize the word to make the function public
	publicFunctionName := strings.ToUpper(string(fn.name[0])) + fn.name[1:]

	textTemplate := fmt.Sprintf("func %s(%s) %s {\n {{.}} \n}\n", publicFunctionName, parameters, fn.returnType)

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
