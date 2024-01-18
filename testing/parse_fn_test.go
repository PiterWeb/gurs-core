package testing

import (
	"fmt"
	"testing"

	gurs_core "github.com/PiterWeb/gurs-core"
)

func getMockRustFunctionsRaw() string {

	return `[{rust\basic_test paramVec [{adios Vec<&str>}] } {rust\basic_test twoParams [{xd str} {lol i32}] } {rust\basic_test unsafeReturn [] &str}]`

}

func TestParseFn(t *testing.T) {

	rustFunctions := gurs_core.GetFunctions(getMockRustFileNames())

	t.Logf("Rust functions parsed: %v", rustFunctions)

	if len(rustFunctions) != RUST_FUNCTIONS {
		t.Fatalf("Number of parsed rust functions is not correct %d/%d", len(rustFunctions), RUST_FUNCTIONS)
	}

	if getMockRustFunctionsRaw() != fmt.Sprintf("%v", rustFunctions) {
		t.Fatalf("The expected result (%s) is not equal to the actual : %v", getMockRustFunctionsRaw(), rustFunctions)
	}

}
