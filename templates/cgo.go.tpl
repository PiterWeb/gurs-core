// File automatically generated with gurs ({{.GursVersion}})
// modify it at your own risk

package {{.Package}}

// #include "{{.Package}}.h"
import "C"

{{range .Functions}}

    {{ReplaceFnBody .ToString `
        // This is a placeholder for the function body
        
    `}}

{{end}}