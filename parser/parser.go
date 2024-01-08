package parser

import (
	"strings"
	"sync"
)

func convertRustFnToStruct(rustFunction string) *Rustfn {

	fnName := strings.Split(rustFunction, `fn`)

	if len(fnName) != 2 {
		return nil
	}

	fnName = strings.Split(fnName[1], "(")

	if len(fnName) != 2 {
		return nil
	}

	return &Rustfn{
		name: fnName[0],
	}
}

func GetFunctions(filePaths []string) []Rustfn {

	var wg sync.WaitGroup

	functions := []Rustfn{}

	for _, path := range filePaths {

		wg.Add(1)

		go func(path string) {

			defer wg.Done()

			functionSignatures := parseRustFile(path)

			for _, signature := range functionSignatures {

				fnStruct := convertRustFnToStruct(signature)

				if fnStruct == nil {
					continue
				}

				functions = append(functions, *fnStruct)

			}

		}(path)

	}

	wg.Wait()

	return functions

}
