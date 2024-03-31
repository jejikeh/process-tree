package filemap

import (
	"fmt"
	"io/fs"
	"path/filepath"

	"github.com/jejikeh/process-tree/treemap"
)

func InitTreemap(path string) (*treemap.Treemap, error) {
	if err := filepath.WalkDir(path, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			fmt.Printf("DIRECTORY: %s\n", d.Name())
		} else {
			fmt.Printf("FILE: %s\n", d.Name())
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return nil, nil
}
