package testing

import (
	"encoding/json"
	"os"
	"reflect"
	"testing"

	gurs_core "github.com/PiterWeb/gurs-core"
)

func TestParseFn(t *testing.T) {

	rustFunctions := gurs_core.GetFunctions(getMockRustFileNames())

	t.Logf("Rust functions parsed: %v", rustFunctions)

	if len(rustFunctions) != RUST_FUNCTIONS {
		t.Fatalf("Number of parsed rust functions is not correct %d/%d", len(rustFunctions), RUST_FUNCTIONS)
	}

	if !reflect.DeepEqual(getMockRustFunctions(), rustFunctions) {
		t.Fatalf("The expected result (%s) is not equal to the actual: (%v)", getMockRustFunctions(), rustFunctions)
	}

	f, err := os.Create("./out/parse_fn.json")

	if err != nil {
		t.Fatalf("Error creating file: %v", err)
	}

	defer f.Close()

	rsFnData, err := json.Marshal(rustFunctions)

	if err != nil {
		t.Fatalf("Error marshalling data: %v", err)
	}

	f.Write(rsFnData)

}
