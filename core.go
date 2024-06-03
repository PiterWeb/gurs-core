// 🦀 Rust code parser & type-safe interface for Golang. ✨
package gurs_core

import (
	"os"
	"path/filepath"

	"github.com/PiterWeb/gurs-core/explore"
	"github.com/PiterWeb/gurs-core/parser"
	"github.com/PiterWeb/gurs-core/templates"
)

const GURS_VERSION = "0.1.0"

type Gofunc = parser.GoFn

type ExecuteOptions struct {
	Destination string
	Pkg         string
}

// GetFunctions parse the files according to the slice of filePaths
// from the arguments.
// It returns an slice of Rustfn(structs) that stores the
// name, parameter (with types) & the returned values (with types)
func GetFunctions(filePaths []string) []parser.RustFn {
	return parser.GetFunctions(filePaths)
}

// ConvertRsFnSliceToGo gets an slice of RustFn and returns the corresponding transpilation to Gofunc slice.
// For individual implementation is a Rustfn struct method  (Rustfn -> Gofunc)
func ConvertRsFnSliceToGo(fns *[]parser.RustFn) (goFuncs []Gofunc) {
	return parser.ConvertRsFnSliceToGo(*fns)
}

// ExploreFolder gets the rootFolder path and returns the filePath of all the .rs files located on the folder
func ExploreFolder(folderPath string) ([]string, error) {

	return explore.ExploreFolder(folderPath)

}

func CGo(goFuncs []Gofunc, options ExecuteOptions) {

	cgoTemp, err := templates.Cgo()

	if err != nil {
		panic(err)
	}

	outputFile, err := os.Create(filepath.Join(options.Destination, options.Pkg+".go"))

	if err != nil {
		panic(err)
	}

	cgoTemp.Execute(outputFile, templates.CgoTemplate{
		Functions:   goFuncs,
		Package:     options.Pkg,
		GursVersion: GURS_VERSION,
	})

}
