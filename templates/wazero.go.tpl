// File automatically generated with gurs ({{.GursVersion}})
// modify it at your own risk

package {{.Package}}

import "{{.Wazero_Runtime}}"

{{range .Functions}}
    {{.Wazero}}
{{end}}