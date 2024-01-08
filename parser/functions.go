package parser

type rustfn struct {
	name       string
	parameters []parameter
	returnType string
}

type parameter struct {
	name      string
	valueType string
}

func (f rustfn) GetName() string {
	return f.name
}

func (f rustfn) GetParameters() []parameter {
	return f.parameters
}

func (f rustfn) GetReturnType() string {
	return f.returnType
}

func (p parameter) GetName() string {
	return p.name
}

func (p parameter) GetType() string {
	return p.valueType
}
