package parser

import (
	"strings"
	"sync"
)

func convertRustFnToStruct(rustFunction string) *rustfn {

	fnName := strings.Split(rustFunction, `fn`)

	if len(fnName) != 2 {
		return nil
	}

	fnName = strings.Split(fnName[1], "(")

	if len(fnName) != 2 {
		return nil
	}

	return &rustfn{
		name: fnName[0],
	}
}

func GetFunctions(filePaths []string) []rustfn {

	var wg sync.WaitGroup

	functions := []rustfn{}

	for _, path := range filePaths {

		wg.Add(1)

		go func(path string) {

			functionSignatures := parseRustFile(path)

			for _, signature := range functionSignatures {

				fnStruct := convertRustFnToStruct(signature)

				if fnStruct == nil {
					continue
				}

				functions = append(functions, *fnStruct)

			}

			wg.Done()

		}(path)

	}

	wg.Wait()

	return functions

}
