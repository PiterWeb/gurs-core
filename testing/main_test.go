package main

import (
	"testing"

	"github.com/PiterWeb/gurs-core/parser"
)

func Test_Main(t *testing.T) {

	rustFunctions := parser.GetFunctions([]string{
		"./test_file.rs",
	})

	t.Log(rustFunctions)

}
