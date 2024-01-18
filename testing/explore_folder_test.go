package testing

import (
	"slices"
	"testing"

	gurs_core "github.com/PiterWeb/gurs-core"
)

func getMockRustFileNames() []string {
	return []string{
		`rust\basic_test.rs`,
	}
}

func TestExploreFolder(t *testing.T) {

	rustFiles, err := gurs_core.ExploreFolder("./rust")

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Founded rust files : %v", rustFiles)

	if len(rustFiles) != RUST_FILES {
		t.Fatalf("Number of found rust files is not correct %d/%d", len(rustFiles), RUST_FILES)
	}

	for _, fn := range rustFiles {
		if b := slices.Contains(getMockRustFileNames(), fn); !b {
			t.Fatalf("Rust file (%s) not found in %v", fn, getMockRustFileNames())
		}
	}

}
