package runtime

import (
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"slices"

	"github.com/PiterWeb/gurs-core/parser"
	"github.com/PiterWeb/gurs-core/templates"
)

type Gofunc = parser.GoFn

func CGoRuntime(goFuncs []Gofunc, options ExecuteOptions) {

	cgoTemp, err := templates.Cgo()

	if err != nil {
		panic(err)
	}

	outputFile, err := os.Create(filepath.Join(options.Destination, options.Pkg+".go"))

	if err != nil {
		panic(err)
	}

	defer outputFile.Close()

	go func() {
		cgoTemp.Execute(outputFile, templates.CgoTemplate{
			Functions:   goFuncs,
			Package:     options.Pkg,
			GursVersion: GURS_VERSION,
		})
	}()

	CompileRust(goFuncs, options)

}

func GetFileNames(fn []Gofunc) []string {

	var fileNames []string

	for _, f := range fn {
		if slices.Contains(fileNames, f.FileName) {
			continue
		}
		fileNames = append(fileNames, f.FileName)
	}

	return fileNames

}

func CompileRust(fn []Gofunc, options ExecuteOptions) {

	fileNames := GetFileNames(fn)

	for _, f := range fileNames {

		srcDir := path.Dir(f)

		cmd := exec.Command("cd .. && cargo build --release --target-dir ")

		cmd.Dir = srcDir

		err := cmd.Run()

		if err != nil {
			panic(err)
		}

	}

}
