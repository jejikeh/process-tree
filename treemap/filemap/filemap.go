package filemap

import (
	"fmt"
	"io/fs"
	"path/filepath"

	"github.com/jejikeh/process-tree/treemap"
)

func InitTreemap(path string) (*treemap.Treemap, error) {
	tree := treemap.NewTreemap()

	if err := filepath.WalkDir(path, func(path string, d fs.DirEntry, err error) error {
		return visitFile(path, &tree, d, err)
	}); err != nil {
		return nil, err
	}

	return nil, nil
}

func visitFile(path string, tree *treemap.Treemap, d fs.DirEntry, err error) error {
	if err != nil {
		return err
	}

	if d.IsDir() {
		fmt.Printf("DIRECTORY: %s\n", filepath.Dir(d.Name()))
	} else {
		fmt.Printf("FILE: %s\n", d.Name())
	}

	return nil
}
