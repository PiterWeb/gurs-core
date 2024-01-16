// gurs-core sub-package that parses: files / functions (fn) / needed structs
package parser

import (
	"fmt"
	"regexp"
	"strings"
	"sync"
)

func convertRustFnToStruct(filePath string, rustFunction string) *Rustfn {

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

	parameters := []parameter{}

	for i := range stringParameters {

		// Trim Spaces & then split by spaces to get name and type of param
		parameterValues := strings.Split(stringParameters[i], ":")

		if len(parameterValues) == 2 {

			parameters = append(parameters, parameter{
				name:      strings.TrimSpace(parameterValues[0]),
				valueType: strings.TrimSpace(parameterValues[1]),
			})
		}

	}

	returnTypeRegex := regexp.MustCompile(`fn\s+[A-Za-z0-9]+[\s?]*\([\s?]*[^)]*[\s?]*\)[\s?]*-?>?[\s?]*([^)]+?)\s{?`)
	returnType := returnTypeRegex.FindStringSubmatch(rustFunction)[1]

	// If no return type regex catch the '{' so we replace it with a blank space
	returnType = strings.ReplaceAll(returnType, "{", "")

	return &Rustfn{
		name:       fnName,
		fileName:   fileName[0],
		parameters: parameters,
		returnType: returnType,
	}
}

func GetFunctions(filePaths []string) []Rustfn {

	var wg sync.WaitGroup

	functions := []Rustfn{}

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
