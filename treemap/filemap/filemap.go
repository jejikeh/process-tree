package filemap

import (
	"fmt"
	"io/fs"
	"path/filepath"
	// "sort"
	"strings"

	"github.com/jejikeh/process-tree/treemap"
)

var foundFolders = make(map[string]*treemap.Node)

func InitTreemap(path string) (treemap.Treemap, error) {
	tree := treemap.NewTreemap()

	if err := filepath.WalkDir(path, func(path string, d fs.DirEntry, err error) error {
		return visitFile(path, &tree, d, err)
	}); err != nil {
		return tree, err
	}

	return tree, nil
}

func visitFile(path string, tree *treemap.Treemap, d fs.DirEntry, err error) error {
	if err != nil {
		return err
	}

	if d.IsDir() {
		if err = addDirectoryNode(tree, path); err != nil {
			return err
		}
	} else {
		fileInfo, err := d.Info()

		if err != nil {
			return err
		}

		if err = addFileNode(tree, path, fileInfo); err != nil {
			return err
		}
	}

	return nil
}

func addDirectoryNode(tree *treemap.Treemap, path string) error {
	path = strings.TrimSuffix(path, "/")

	basepath := filepath.Dir(path)

	var parent *treemap.Node

	parent, parentWasFound := foundFolders[basepath]

	if !parentWasFound && len(tree.Nodes) > 0 {
		return fmt.Errorf("parent not found for '%s', root '%s'\n", basepath, tree.Nodes[0].Name)
	}

	node, err := tree.Add(path, parent)

	if err != nil {
		return err
	}

	foundFolders[path] = node

	return nil
}

func addFileNode(tree *treemap.Treemap, path string, fileInfo fs.FileInfo) error {
	var parent *treemap.Node

	basepath := filepath.Dir(path)

	parent, parentWasFound := foundFolders[basepath]

	if !parentWasFound {
		return fmt.Errorf("the file %s has unknown directory!", basepath)
	}

	node, err := tree.Add(path, parent)

	if err != nil {
		return err
	}

	node.Size = float64(fileInfo.Size())

	return nil
}
