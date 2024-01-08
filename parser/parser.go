// gurs-core sub-package that parses: files / functions (fn) / needed structs
package parser

import (
	"fmt"
	"strings"
	"sync"
)

func convertRustFnToStruct(filePath string, rustFunction string) *Rustfn {

	fnName := strings.Split(rustFunction, `fn`)

	if len(fnName) != 2 {
		return nil
	}

	fnName = strings.Split(fnName[1], "(")

	if len(fnName) != 2 {
		return nil
	}

	fileName := strings.Split(filePath, ".rs")

	if len(fnName) != 2 {
		return nil
	}

	stringParameters := strings.Split(fnName[1], ")")

	if len(stringParameters) < 1 {
		return nil
	}

	stringParameters = strings.Split(stringParameters[0], ",")

	parameters := []parameter{}

	for i := range stringParameters {

		// Trim Spaces & then split by spaces to get name and type of param
		parameterValues := strings.Split(strings.TrimSpace(stringParameters[i]), " ")

		if len(parameterValues) == 2 {

			parameters = append(parameters, parameter{
				name:      parameterValues[0],
				valueType: parameterValues[1],
			})
		}

	}

	return &Rustfn{
		name:       fnName[0],
		fileName:   fileName[0],
		parameters: parameters,
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

				fnStruct := convertRustFnToStruct(path, signature)

				if fnStruct == nil {
					fmt.Printf("An error ocurred during fn parsing for : %s", signature)
					continue
				}

				functions = append(functions, *fnStruct)

			}

		}(path)

	}

	wg.Wait()

	return functions

}
