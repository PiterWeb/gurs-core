package translate

var translationTypes = map[string]string{
	"i8":      "int8",
	"i16":     "int16",
	"i32":     "int32",
	"i64":     "int64",
	"u8":      "uint8",
	"u16":     "uint16",
	"u32":     "uint32",
	"u64":     "uint64",
	"f32":     "float32",
	"f64":     "float64",
	"boolean": "bool",
	"char":    "rune",
	"str":     "string",
	"&str":    "string",
	"enum":    "interface{}",
	"Vec":     "[]",
}
