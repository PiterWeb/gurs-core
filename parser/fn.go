package parser

import "fmt"

type Rustfn struct {
	fileName   string
	name       string
	parameters []parameter
	returnType string
}

type Gofunc Rustfn

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
