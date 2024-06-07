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

var goToCGoArgTypes = map[string]string{
	"string":  "*C.char",
	"byte":    "C.uchar", // Unsigned char for byte
	"rune":    "C.int",   // Rune is a Unicode code point, equivalent to int32
	"int":     "C.int",
	"int8":    "C.schar",
	"int16":   "C.short",
	"int32":   "C.int",
	"int64":   "C.slong",
	"uint":    "C.uint",
	"uint8":   "C.uchar",
	"uint16":  "C.ushort",
	"uint32":  "C.uint",
	"uint64":  "C.ulong",
	"float32": "C.float",
	"float64": "C.double",
	"bool":    "C.bool", // bool in Go corresponds to _Bool in C
}

var conversionFunctionsGoToCGo = map[string]string{
	"string": "C.CString",
}

var conversionFunctionsCGoToGo = map[string]string{
	"string": "C.GoString",
}

func goTypeToCGoType(goType string) string {
	// Regular expression to match array types (e.g., [5]int, []int, [][5]int)
	arrayRegex := regexp.MustCompile(`\[(.*?)\](.*)`)

	matches := arrayRegex.FindStringSubmatch(goType)

	// Check if it's an array type
	if len(matches) == 3 {
		size := matches[1] // array size, could be empty for slices
		baseType := strings.TrimSpace(matches[2])
		cgoBaseType, ok := goToCGoArgTypes[baseType]
		if !ok {
			// If the base type is not found, return interface{} (any type in Cgo)
			return "interface{}"
		}
		if size == "" {
			return fmt.Sprintf("*%s", cgoBaseType) // for slices, use pointer type
		}
		return fmt.Sprintf("[%s]%s", size, cgoBaseType)
	}

	// Not an array type, directly map the base type
	cgoBaseType, ok := goToCGoArgTypes[goType]
	if !ok {
		// If the base type is not found, return interface{} (any type in Cgo)
		return "interface{}"
	}
	return cgoBaseType
}

func GoFuncSignatureToCGo(fn _ICGoFunc) string {

	fmt.Println("GoFuncSignatureToCGo")

	// Convert the parameters to a string
	parameters := strings.Join(fn.GetRawParameters(), ", ")

	// Replace the colon with a space
	parameters = strings.ReplaceAll(parameters, ":", " ")

	returnTypeToCGo := goTypeToCGoType(fn.GetReturnType())

	textTemplate := fmt.Sprintf("func %s(%s) %s { {{.}} }\n", fn.GetName(), parameters, returnTypeToCGo)

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

		cgo_conversion, ok := conversionFunctionsGoToCGo[p.GetType()]

		if !ok {

			if strings.Contains(p.GetType(), "[]") {

				string_slice := strings.Split(p.GetType(), "[]")

				pure_type := string_slice[len(string_slice)-1]

				cgo_conversion, ok = conversionFunctionsGoToCGo[pure_type]

				if !ok {
					// If the conversion function is not found, use the default conversion
					cgo_conversion = goTypeToCGoType(p.GetType())
				} else {
					// If it's a slice, use a pointer type
					cgo_conversion = fmt.Sprintf("*%s", cgo_conversion)
				}

			} else {
				// If the conversion function is not found, use the default conversion
				cgo_conversion = goTypeToCGoType(p.GetType())
			}

		}

		arg_name := p.GetName()
		variables += fmt.Sprintf("c_%s := %s(%s)\n", arg_name, cgo_conversion, arg_name)
	}

	return variables
}
