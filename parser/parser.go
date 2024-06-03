// gurs-core sub-package that parses: files / functions (fn) / needed structs
package parser

import (
	"fmt"
	"regexp"
	"strings"
	"sync"

	"github.com/PiterWeb/gurs-core/functions"
)

func convertRustFnToStruct(filePath string, rustFunction string) *RustFn {

	fnNameRegex := regexp.MustCompile(`fn\s+([A-Za-z0-9]+)+[\s?]*\(`)

	fnName := fnNameRegex.FindStringSubmatch(rustFunction)[1]

	if fnName == "" {
		return nil
	}

	fileName := strings.Split(filePath, ".rs")

	if len(fileName) != 2 {
		return nil
	}

	stringParametersRegex := regexp.MustCompile(`fn\s+[A-Za-z0-9]+[\s?]*\([\s?]*([^)]+)*[\s?]*\)`)

	stringParameters := strings.Split(stringParametersRegex.FindStringSubmatch(rustFunction)[1], ",")

	parameters := []functions.Parameter{}

	for i := range stringParameters {

		// Trim Spaces & then split by spaces to get name and type of param
		parameterValues := strings.Split(stringParameters[i], ":")

		if len(parameterValues) == 2 {

			parameters = append(parameters, functions.Parameter{
				Name:      strings.TrimSpace(parameterValues[0]),
				ValueType: strings.TrimSpace(parameterValues[1]),
			})
		}

	}

	returnTypeRegex := regexp.MustCompile(`fn\s+[A-Za-z0-9]+[\s?]*\([\s?]*[^)]*[\s?]*\)[\s?]*-?>?[\s?]*([^)]+?)\s{?`)
	returnType := returnTypeRegex.FindStringSubmatch(rustFunction)[1]

	// If no return type regex catch the '{' so we replace it with a blank space
	returnType = strings.ReplaceAll(returnType, "{", "")

	return &RustFn{
		Name:       fnName,
		FileName:   fileName[0],
		Parameters: parameters,
		ReturnType: returnType,
	}
}

func GetFunctions(filePaths []string) []RustFn {

	var wg sync.WaitGroup

	functions := []RustFn{}

	for _, path := range filePaths {

		wg.Add(1)

		go func(path string) {

			defer wg.Done()

			functionSignatures, err := parseRustFile(path)

			if err != nil {
				fmt.Printf("An error ocurred during file parsing for: %s", path)
				return
			}

			for _, signature := range functionSignatures {

				fnStruct := convertRustFnToStruct(path, signature)

				if fnStruct == nil {
					fmt.Printf("An error ocurred during fn parsing for:	%s::%s", path, signature)
					continue
				}

				functions = append(functions, *fnStruct)

			}

		}(path)

	}

	wg.Wait()

	return functions

}
