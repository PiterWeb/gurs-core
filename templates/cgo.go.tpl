// File automatically generated with gurs ({{.GursVersion}})
// modify it at your own risk

package {{.Package}}

// #cgo CFLAGS: -g -Wall
// #include <stdlib.h>
// #include "{{.Package}}.h"
import "C"

{{range .Functions}}

    {{$varNames := GetVariableNames .}}
    {{$fnCall := printf "return C.%s(%v)" .Name $varNames}}
    {{$varDeclarations := CreateVariables .}}

    {{ReplaceFnBody .ToString $varDeclarations $fnCall}}

{{end}}