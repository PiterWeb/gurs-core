package testing

import (
	"testing"

	gurs_core "github.com/PiterWeb/gurs-core"
)

func Test_Main(t *testing.T) {

	rustFunctions := gurs_core.GetFunctions([]string{
		"./test_file.rs",
	})

	t.Log(rustFunctions)

	for _, fn := range rustFunctions {

		goTypes := gurs_core.TranspileTypes(fn.GetRawParameters())

		goTypes = append(goTypes, gurs_core.TranspileTypes([]string{fn.GetReturnType()})[0])

		t.Log(goTypes)
	}

}
