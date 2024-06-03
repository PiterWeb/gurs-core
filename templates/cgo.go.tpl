// File automatically generated with gurs ({{.GursVersion}})
// modify it at your own risk

package {{.Package}}

// #cgo CFLAGS: -g -Wall
// #include <stdlib.h>
// #include "{{.Package}}.h"
import "C"

{{range .Functions}}

    {{.ToCGo}}

{{end}}