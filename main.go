package main

import (
	"fmt"

	"github.com/PiterWeb/gurs-core/parser"
)

func main() {

	rustFunctions := parser.GetFunctions()

	fmt.Println(rustFunctions)

}
