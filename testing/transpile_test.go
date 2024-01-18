package testing

import (
	"reflect"
	"testing"

	gurs_core "github.com/PiterWeb/gurs-core"
)

func TestTranspile(t *testing.T) {

	rustFunctions := getMockRustFunctions()

	goFunctions := gurs_core.ConvertRsFnSliceToGo(&rustFunctions)

	t.Logf("Go compatible functions (%d):  \n %s", len(goFunctions), goFunctions)

	if len(goFunctions) != GO_COMPATIBLE_FUNCTIONS {
		t.Fatalf("Number of Golang functions is not correct: %d/%d", len(goFunctions), GO_COMPATIBLE_FUNCTIONS)
	}

	if !reflect.DeepEqual(goFunctions, getMockGoFunctions()) {
		t.Fatalf("The expected result (%s) is not equal to the actual: (%v)", getMockGoFunctions(), goFunctions)

	}

}
