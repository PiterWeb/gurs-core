package testing

import (
	"testing"

	gurs_core "github.com/PiterWeb/gurs-core"
)

func Test_Main(t *testing.T) {

	rustFiles, err := gurs_core.ExploreFolder(".")
	t.Run("Explore Folder", func(t *testing.T) {

		if err != nil {
			t.Fatal(err)
		}

		t.Logf("Founded rust files : %v", rustFiles)

	})

	rustFunctions := gurs_core.GetFunctions(rustFiles)

	t.Run("Parse Rust functions", func(t *testing.T) {

		t.Logf("Rust functions parsed: %v", rustFunctions)

		if len(rustFunctions) != 3 {
			t.Fatal("Number of parsed rust functions is not correct")
		}

	})

	goFunctions := gurs_core.ConvertRsFnSliceToGo(&rustFunctions)

	t.Run("Transpilation to Golang", func(t *testing.T) {

		t.Logf("Go compatible functions: %v", goFunctions)

		if len(goFunctions) != 3 {
			t.Logf("Number of Golang functions is not correct ")
		}

	})

}
