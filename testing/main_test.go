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

}
