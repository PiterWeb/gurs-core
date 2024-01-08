// gurs-core - run rust 🦀 code in your golang projects ✨
package gurs_core

import (
	"github.com/PiterWeb/gurs-core/parser"
	"github.com/PiterWeb/gurs-core/transpile"
)

// GetFunctions parse the files according to the slice of filePaths
// from the arguments.
// It returns an slice of Rustfn(structs) that stores the
// name, parameter (with types) & the returned values (with types)
func GetFunctions(filePaths []string) []parser.Rustfn {
	return parser.GetFunctions(filePaths)
}

// TranspileTypes gets the rust types and transpiles to golang valid types
func TranspileTypes(rustTypes []string) []string {
	return transpile.TranspileTypes(rustTypes)
}
