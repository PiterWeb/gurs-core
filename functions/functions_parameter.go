package functions

type Iparameter interface {
	GetName() string
	GetType() string
}

type Parameter struct {
	Name      string `json:"name"`
	ValueType string `json:"valueType"`
}

func (p Parameter) GetName() string {
	return p.Name
}

func (p Parameter) GetType() string {
	return p.ValueType
}
