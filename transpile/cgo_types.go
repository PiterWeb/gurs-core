package transpile

import (
	"fmt"
	"regexp"
	"strings"
)

var GoToCGoTypes = map[string]string{
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

func GoTypeToCGoType(goType string) string {
	// Regular expression to match array types (e.g., [5]int, []int, [][5]int)
	arrayRegex := regexp.MustCompile(`\[(.*?)\](.*)`)

	matches := arrayRegex.FindStringSubmatch(goType)
	if len(matches) == 3 {
		size := matches[1] // array size, could be empty for slices
		baseType := strings.TrimSpace(matches[2])
		cgoBaseType, ok := GoToCGoTypes[baseType]
		if !ok {
			return "unknown"
		}
		if size == "" {
			return fmt.Sprintf("*%s", cgoBaseType) // for slices, use pointer type
		}
		return fmt.Sprintf("[%s]%s", size, cgoBaseType)
	}

	// Not an array type, directly map the base type
	cgoBaseType, ok := GoToCGoTypes[goType]
	if !ok {
		return "unknown"
	}
	return cgoBaseType
}
