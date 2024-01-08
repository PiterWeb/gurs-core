package main

import (
	"github.com/PiterWeb/gurs-core/parser"
)

func GetFunctions(filePaths []string) []parser.Rustfn {
	return parser.GetFunctions(filePaths)
}
