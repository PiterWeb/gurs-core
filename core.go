// gurs-core - run rust 🦀 code in your golang projects ✨
package gurs_core

import (
	"github.com/PiterWeb/gurs-core/parser"
)

type Gofunc = parser.Gofunc

// GetFunctions parse the files according to the slice of filePaths
// from the arguments.
// It returns an slice of Rustfn(structs) that stores the
// name, parameter (with types) & the returned values (with types)
func GetFunctions(filePaths []string) []parser.Rustfn {
	return parser.GetFunctions(filePaths)
}

// ConvertRsFnSliceToGo gets an slice of RustFn and returns the corresponding transpilation to Gofunc slice
func ConvertRsFnSliceToGo(fns *[]parser.Rustfn) (goFuncs []Gofunc) {
	return parser.ConvertRsFnSliceToGo(*fns)
}
