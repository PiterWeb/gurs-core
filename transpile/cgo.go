package transpile

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/PiterWeb/gurs-core/functions"
)

// _ICGoFunc is an interface for Go functions that can be converted to Cgo functions
type _ICGoFunc interface {
	GetRawParameters() []string
	GetName() string
	GetReturnType() string
	GetParameters() []functions.Parameter
}

var goToCGoTypes = map[string]string{
	"string":  "C.CString",
	"byte":    "C.uchar",   // Unsigned char for byte
	"rune":    "C.int32_t", // Rune is a Unicode code point, equivalent to int32
	"int":     "C.int",
	"int8":    "C.int8_t",
	"int16":   "C.int16_t",
	"int32":   "C.int32_t",
	"int64":   "C.int64_t",
	"uint":    "C.uint",
	"uint8":   "C.uint8_t",
	"uint16":  "C.uint16_t",
	"uint32":  "C.uint32_t",
	"uint64":  "C.uint64_t",
	"float32": "C.float",
	"float64": "C.double",
	"bool":    "C._Bool", // bool in Go corresponds to _Bool in C
}

func goTypeToCGoType(goType string) string {
	// Regular expression to match array types (e.g., [5]int, []int, [][5]int)
	arrayRegex := regexp.MustCompile(`\[(.*?)\](.*)`)

	matches := arrayRegex.FindStringSubmatch(goType)
	if len(matches) == 3 {
		size := matches[1] // array size, could be empty for slices
		baseType := strings.TrimSpace(matches[2])
		cgoBaseType, ok := goToCGoTypes[baseType]
		if !ok {
			return "unknown"
		}
		if size == "" {
			return fmt.Sprintf("*%s", cgoBaseType) // for slices, use pointer type
		}
		return fmt.Sprintf("[%s]%s", size, cgoBaseType)
	}

	// Not an array type, directly map the base type
	cgoBaseType, ok := goToCGoTypes[goType]
	if !ok {
		return "unknown"
	}
	return cgoBaseType
}

func GoFuncSignatureToCGo(fn _ICGoFunc) string {

	fmt.Println("GoFuncSignatureToCGo")

	// Convert the parameters to a string
	parameters := strings.Join(fn.GetRawParameters(), ", ")

	// Replace the colon with a space
	parameters = strings.ReplaceAll(parameters, ":", " ")

	textTemplate := fmt.Sprintf("func %s(%s) %s { {{.}} }\n", fn.GetName(), parameters, fn.GetReturnType())

	fmt.Println("textTemplate: ", textTemplate)

	argsForCall := getArgsForCall(fn)

	fmt.Println("argsForCall: ", argsForCall)

	cgoVariables := getCGoVariables(fn)

	fmt.Println("cgoVariables: ", cgoVariables)

	// The function call in the cgo file
	fnCall := fmt.Sprintf("\nreturn C.%s(%s)\n", fn.GetName(), argsForCall)

	fmt.Println("textTemplate: ", textTemplate)

	// Replace the function body with the cgo variables and the function call
	return strings.Replace(textTemplate, "{{.}}", cgoVariables+fnCall, 1)

}

func getArgsForCall(fn _ICGoFunc) string {

	variables := ""

	for i, p := range fn.GetParameters() {
		// If it's the last parameter, don't add a comma
		if i == len(fn.GetParameters())-1 {
			variables += fmt.Sprintf("c_%s", p.GetName())
			break
		}
		variables += fmt.Sprintf("c_%s, ", p.GetName())
	}

	return variables

}

func getCGoVariables(fn _ICGoFunc) string {

	variables := "\n"

	for _, p := range fn.GetParameters() {
		cgo_type := goTypeToCGoType(p.GetType())
		arg_name := p.GetName()
		variables += fmt.Sprintf("c_%s := %s(%s)\n", arg_name, cgo_type, arg_name)
	}

	return variables
}
