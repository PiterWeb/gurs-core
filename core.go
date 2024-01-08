// gurs-core - run rust 🦀 code in your golang projects ✨
package gurs_core

import (
	"github.com/PiterWeb/gurs-core/parser"
)

// GetFunctions parse the files according to the slice of filePaths
// from the arguments.
// It returns an slice of Rustfn(structs) that stores the
// name, parameter (with types) & the returned values (with types)
func GetFunctions(filePaths []string) []parser.Rustfn {
	return parser.GetFunctions(filePaths)
}
