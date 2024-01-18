package testing

import (
	"encoding/json"

	"github.com/PiterWeb/gurs-core/parser"
)

const (
	RUST_FILES              = 1
	RUST_FUNCTIONS          = 3
	GO_COMPATIBLE_FUNCTIONS = 3
)

const MOCK_RUST_FN = `[{"fileName":"rust\\basic_test","name":"paramVec","parameters":[{"name":"adios","valueType":"Vec\u003c\u0026str\u003e"}],"returnType":""},{"fileName":"rust\\basic_test","name":"twoParams","parameters":[{"name":"xd","valueType":"str"},{"name":"lol","valueType":"i32"}],"returnType":""},{"fileName":"rust\\basic_test","name":"unsafeReturn","parameters":[],"returnType":"\u0026str"}]`

const MOCK_GO_FN = `[{"fileName":"rust\\basic_test","name":"ParamVec","parameters":[{"name":"adios","valueType":"[]string"}],"returnType":""},{"fileName":"rust\\basic_test","name":"TwoParams","parameters":[{"name":"xd","valueType":"string"},{"name":"lol","valueType":"int32"}],"returnType":""},{"fileName":"rust\\basic_test","name":"UnsafeReturn","parameters":[],"returnType":"string"}]`

func getMockRustFunctions() []parser.Rustfn {

	rustFn := new([]parser.Rustfn)

	json.Unmarshal([]byte(MOCK_RUST_FN), rustFn)

	return *rustFn

}

func getMockGoFunctions() []parser.Gofunc {

	goFn := new([]parser.Gofunc)

	json.Unmarshal([]byte(MOCK_GO_FN), goFn)

	return *goFn

}
