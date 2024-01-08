package parser

import (
	"strings"
)

func convertRustFnToStruct(rustFunction string) *rustfn {

	fnName := strings.Split(rustFunction, `fn`)

	if len(fnName) != 2 {
		return nil
	}

	return &rustfn{
		name: fnName[1],
	}
}

func GetFunctions() []rustfn {

	functionSignatures := parseRustFile("./parser/test_file.rs")

	functions := []rustfn{}

	for _, signature := range functionSignatures {

		fnStruct := convertRustFnToStruct(signature)

		if fnStruct == nil {
			continue
		}

		functions = append(functions, *fnStruct)

	}

	return functions

}
