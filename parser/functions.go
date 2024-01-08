package parser

type Rustfn struct {
	name       string
	parameters []parameter
	returnType string
}

type parameter struct {
	name      string
	valueType string
}

func (f Rustfn) GetName() string {
	return f.name
}

func (f Rustfn) GetParameters() []parameter {
	return f.parameters
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
