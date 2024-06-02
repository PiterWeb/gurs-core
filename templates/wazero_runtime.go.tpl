// File automatically generated with gurs ({{.GursVersion}})
// modify it at your own risk

package wazero_runtime

import (
	"context"

	"github.com/tetratelabs/wazero"
)

var ctx = context.Background()
var r = wazero.NewRuntine(ctx)

func GetRuntime(fn func) wazero.Runtime {
	return r
}

func CloseRuntime() error {
	return r.Close(ctx)
}
