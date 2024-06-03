package parser

import (
	"fmt"
	"strings"
	"text/template"

	"github.com/PiterWeb/gurs-core/functions"
	"github.com/PiterWeb/gurs-core/transpile"
)

type baseFn struct {
	FileName   string                `json:"fileName"`
	Name       string                `json:"name"`
	Parameters []functions.Parameter `json:"parameters"`
	ReturnType string                `json:"returnType"`
}

type RustFn baseFn

func (f baseFn) GetFileName() string {
	return f.FileName
}

func (f baseFn) GetName() string {
	return f.Name
}

func (f baseFn) GetParameters() []functions.Parameter {
	return f.Parameters
}

func (f baseFn) GetRawParameters() []string {

	rawParameters := []string{}

	for _, p := range f.Parameters {
		rawParameters = append(rawParameters, fmt.Sprintf("%s %s", p.GetName(), p.GetType()))
	}

	return rawParameters
}

func (f baseFn) GetReturnType() string {
	return f.ReturnType
}

type GoFn baseFn

func (fn RustFn) ConvertToGo() GoFn {

	convertedParameters := []functions.Parameter{}

	for _, p := range baseFn(fn).GetParameters() {

		convertedParameters = append(convertedParameters, functions.Parameter{
			Name:      p.GetName(),
			ValueType: transpile.TranspileType(p.GetType()),
		})

	}

	// Capitalize the word to make the function public
	publicFunctionName := strings.ToUpper(string(fn.Name[0])) + fn.Name[1:]

	f := GoFn{
		FileName:   fn.FileName,
		Name:       publicFunctionName,
		Parameters: convertedParameters,
		ReturnType: transpile.TranspileType(fn.ReturnType),
	}

	return f

}

func (fn GoFn) ToTemplate() (*template.Template, error) {

	templ := template.New("func-" + fn.Name)

	// Convert the parameters to a string
	parameters := strings.Join(baseFn(fn).GetRawParameters(), ", ")

	// Replace the colon with a space
	parameters = strings.ReplaceAll(parameters, ":", " ")

	textTemplate := fmt.Sprintf("func %s(%s) %s { {{.}} }\n", fn.Name, parameters, fn.ReturnType)

	fmt.Println("textTemplate: ", textTemplate)

	templ, err := templ.Parse(textTemplate)

	if err != nil {
		return nil, err
	}

	return templ, nil

}

func (fn GoFn) ToCGo() string {

	baseFn := baseFn(fn)

	return transpile.GoFuncSignatureToCGo(baseFn)
}

func ConvertRsFnSliceToGo(fns []RustFn) (goFuncs []GoFn) {

	for _, r := range fns {

		goFuncs = append(goFuncs, r.ConvertToGo())

	}

	return

}
