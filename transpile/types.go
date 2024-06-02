package transpile

import "strings"

// common types in rust and their golang equivalents
var transpilationTypes = map[string]string{
	"i8":           "int8",
	"i16":          "int16",
	"i32":          "int32",
	"i64":          "int64",
	"u8":           "uint8",
	"u16":          "uint16",
	"u32":          "uint32",
	"u64":          "uint64",
	"f32":          "float32",
	"f64":          "float64",
	"boolean":      "bool",
	"char":         "rune",
	"str":          "string",
	"String":       "string",
	"&'static str": "string",
}

// Converts Vec<T> to []T
func transpilateSlice(rustVec string) string {

	return "[]" + strings.Split(strings.Split(rustVec, "Vec<")[1], ">")[0]

}

// Transpile Rust type to golang valid type
func TranspileType(rustType string) string {

	isPrimitive := false

	goType := ""

	for transpilationType, goEquivalentType := range transpilationTypes {

		// If type not match on this iteration
		if !strings.Contains(rustType, transpilationType) {
			continue
		}

		// Borrowed type
		borrowedTranspilationType := "&" + transpilationType

		isPrimitive = true

		// If borrowed type
		if strings.Contains(rustType, borrowedTranspilationType) {
			goType = strings.ReplaceAll(rustType, borrowedTranspilationType, goEquivalentType)
			break
		}

		// If normal type
		goType = strings.ReplaceAll(rustType, transpilationType, goEquivalentType)
		break

	}

	// If no match with primitive types
	if !isPrimitive {
		return transpileCustomType(rustType)
	}

	// If type is a Vec<?>
	if strings.Contains(goType, "Vec<") {
		goType = transpilateSlice(goType)
	}

	return goType

}

// Transpile array of rust types to golang valid types
func TranspileTypes(rustTypes []string) []string {

	goTypes := []string{}

	for _, rustType := range rustTypes {

		goType := TranspileType(rustType)

		goTypes = append(goTypes, goType)
	}

	return goTypes

}
