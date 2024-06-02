package testing

import (
	"os"
	"reflect"
	"testing"

	gurs_core "github.com/PiterWeb/gurs-core"
)

func TestTranspile(t *testing.T) {

	rustFunctions := getMockRustFunctions()

	goFunctions := gurs_core.ConvertRsFnSliceToGo(&rustFunctions)

	t.Logf("Rust functions (%d):  \n %s", len(rustFunctions), rustFunctions)

	t.Logf("Go compatible functions (%d):  \n %s", len(goFunctions), goFunctions)

	if len(goFunctions) != GO_COMPATIBLE_FUNCTIONS {
		t.Fatalf("Number of Golang functions is not correct: %d/%d", len(goFunctions), GO_COMPATIBLE_FUNCTIONS)
	}

	if !reflect.DeepEqual(goFunctions, getMockGoFunctions()) {
		t.Fatalf("The expected result (%s) is not equal to the actual: (%v)", getMockGoFunctions(), goFunctions)

	}

	file, err := os.Create("./out/transpile.go")

	if err != nil {
		t.Fatalf("Error creating file: %v", err)
	}

	defer file.Close()

	n, err := file.WriteString("package testing\n\n")

	if err != nil {
		t.Fatalf("Error writing to file: %v", err)
	}

	t.Logf("Wrote %d bytes to file", n)

	for _, fn := range goFunctions {
		temp, err := fn.ToTemplate()
		if err != nil {
			t.Fatalf("Error converting function to string: %v", err)
		}

		// Write to file and pass empty string as data for the body of the function
		temp.Execute(file, "")

	}

}
