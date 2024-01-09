package parser

import (
	"fmt"

	"github.com/PiterWeb/gurs-core/transpile"
)

type Rustfn struct {
	fileName   string
	name       string
	parameters []parameter
	returnType string
}

type parameter struct {
	name      string
	valueType string
}

func (f Rustfn) GetFileName() string {
	return f.fileName
}

func (f Rustfn) GetName() string {
	return f.name
}

func (f Rustfn) GetParameters() []parameter {
	return f.parameters
}

func (f Rustfn) GetRawParameters() []string {

	rawParameters := []string{}

	for _, p := range f.parameters {
		rawParameters = append(rawParameters, fmt.Sprintf("%s %s", p.name, p.valueType))
	}

	return rawParameters
}

func (f Rustfn) GetReturnType() string {
	return f.returnType
}

func (p parameter) GetName() string {
	return p.name
}

func (p parameter) GetType() string {
	return p.valueType
}

type Gofunc Rustfn

func (fn Rustfn) ConvertToGo() Gofunc {

	convertedParameters := []parameter{}

	for _, p := range fn.GetParameters() {

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

func ConvertRsFnSliceToGo(fns []Rustfn) (goFuncs []Gofunc) {

	for _, r := range fns {

		goFuncs = append(goFuncs, r.ConvertToGo())

	}

	return

}
