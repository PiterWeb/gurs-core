// gurs-core sub-package that manages all the logic binded to the generation of code based on the runtime
package runtime

type runtimet uint

const (
	CGORuntime runtimet = iota
	WazeroRuntime
)

func SelectRuntime(m runtimet) {

	if m == CGORuntime {
		cGoRuntime()
	} else if m == WazeroRuntime {
		wazeroRuntime()
	}

}

func GetRuntimes() map[string]runtimet {
	return map[string]runtimet{
		"CGoRuntime":    CGORuntime,
		"WazeroRuntime": WazeroRuntime,
	}
}
