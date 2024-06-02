package testing

import (
	"encoding/json"
	"os"

	"github.com/PiterWeb/gurs-core/parser"
)

const (
	RUST_FILES              = 1
	RUST_FUNCTIONS          = 4
	GO_COMPATIBLE_FUNCTIONS = 4
)

const MOCK_RUST_FN_FILE = "./assets/parse_fn_rust.json"

const MOCK_GO_FN_FILE = "./assets/parse_fn_go.json"

func getMockRustFunctions() []parser.Rustfn {

	rustFn := new([]parser.Rustfn)

	MOCK_RUST_FN, err := os.ReadFile(MOCK_RUST_FN_FILE)

	if err != nil {
		panic(err)
	}

	json.Unmarshal(MOCK_RUST_FN, rustFn)

	return *rustFn

}

func getMockGoFunctions() []parser.Gofunc {

	goFn := new([]parser.Gofunc)

	MOCK_GO_FN, err := os.ReadFile(MOCK_GO_FN_FILE)

	if err != nil {
		panic(err)
	}

	json.Unmarshal(MOCK_GO_FN, goFn)

	return *goFn

}
