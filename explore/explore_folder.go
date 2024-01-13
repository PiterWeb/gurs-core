package explore

import (
	"io/fs"
	"path/filepath"
	"strings"
)

func ExploreFolder(folderPath string) ([]string, error) {

	filePaths := []string{}

	err := filepath.Walk(folderPath, func(path string, info fs.FileInfo, err error) error {

		if err != nil {
			return err
		}

		isRustFile := strings.Contains(info.Name(), ".rs")

		if info.IsDir() || !isRustFile {
			return nil
		}

		filePaths = append(filePaths, path)

		return nil

	})

	return filePaths, err

}
