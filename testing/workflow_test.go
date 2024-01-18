package testing

import (
	"os"
	"testing"

	gurs_core "github.com/PiterWeb/gurs-core"
)

const (
	RUST_FILES              = 1
	RUST_FUNCTIONS          = 3
	GO_COMPATIBLE_FUNCTIONS = 3
)

func TestWorkflow(t *testing.T) {

	rustFiles, err := gurs_core.ExploreFolder(".")
	t.Run("Explore Folder", func(t *testing.T) {

		if err != nil {
			t.Fatal(err)
		}

		t.Logf("Founded rust files : %v", rustFiles)

		if len(rustFiles) != RUST_FILES {
			t.Fatalf("Number of found rust files is not correct %d/%d", len(rustFiles), RUST_FILES)
		}

	})

	rustFunctions := gurs_core.GetFunctions(rustFiles)

	t.Run("Parse Rust functions", func(t *testing.T) {

		t.Logf("Rust functions parsed: %v", rustFunctions)

		if len(rustFunctions) != RUST_FUNCTIONS {
			t.Fatalf("Number of parsed rust functions is not correct %d/%d", len(rustFunctions), RUST_FUNCTIONS)
		}

	})

	goFunctions := gurs_core.ConvertRsFnSliceToGo(&rustFunctions)

	t.Run("Transpilation to Golang", func(t *testing.T) {

		t.Logf("Go compatible functions: %v", goFunctions)

		if len(goFunctions) != GO_COMPATIBLE_FUNCTIONS {
			t.Fatalf("Number of Golang functions is not correct: %d/%d", len(goFunctions), GO_COMPATIBLE_FUNCTIONS)
		}

	})

	t.Run("Generate Go Templates", func(t *testing.T) {

		for _, fn := range goFunctions {
			templ, err := fn.ToStringTemplate()

			if err != nil {
				t.Fatalf(err.Error())
			}

			templ.Execute(os.Stdout, "// Body template")
		}

	})

}
