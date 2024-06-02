package testing

import (
	"reflect"
	"testing"

	gurs_core "github.com/PiterWeb/gurs-core"
)

func getMockRustFileNames() []string {
	return []string{
		`assets\basic_c_test.rs`,
	}
}

func TestExploreFolder(t *testing.T) {

	rustFiles, err := gurs_core.ExploreFolder("./assets")

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Founded rust files : %v", rustFiles)

	if len(rustFiles) != RUST_FILES {
		t.Fatalf("Number of found rust files is not correct %d/%d", len(rustFiles), RUST_FILES)
	}

	if !reflect.DeepEqual(getMockRustFileNames(), rustFiles) {
		t.Fatalf("The expected result (%s) is not equal to the actual: (%s)", getMockRustFileNames(), rustFiles)
	}

}
