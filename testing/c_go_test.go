package testing

import (
	"os"
	"testing"

	"github.com/PiterWeb/gurs-core/templates"
)

func TestCGo(t *testing.T) {

	goFuncs := getMockGoFunctions()

	file, err := os.Create("./out/cgo.go")

	if err != nil {
		t.Fatalf("Error creating file: %v", err)
	}

	defer file.Close()

	cGoTemp, err := templates.Cgo()

	if err != nil {
		t.Fatalf("Error creating template: %v", err)
	}

	cGoTemp.Execute(file, templates.CgoTemplate{
		Functions:   goFuncs,
		Package:     "cgo",
		GursVersion: "0.1.0",
	})

}
