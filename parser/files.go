package parser

import (
	"fmt"
	"os"
	"regexp"
)

func parseRustFile(fileName string) []string {

	rustFile, err := os.ReadFile(fileName)

	if err != nil {

		fmt.Printf("Rust file %s has not been readed correctly cause -> %s\n", fileName, err.Error())

		return nil
	}

	regexFunctionsFirm := regexp.MustCompile(`#\[no_mangle\]\s+pub\s+u?n?s?a?f?e?[\s?]*extern\s+"C"\s+fn\s+[A-Za-z0-9]+\([^)]*\)`)
	// #\[no_mangle\]\s+pub\s+u?n?s?a?f?e?[\s?]*extern\s+"C"\s+fn\s+[A-Za-z0-9]+\([^)]*\).*[\s?]*-?>?[\s?]*&?[A-Za-z0-9]+.*$

	functionMatches := regexFunctionsFirm.FindAllString(string(rustFile), -1)

	if functionMatches == nil {
		fmt.Printf("No exported functions found on file %s\n", fileName)
	}

	return functionMatches

}
