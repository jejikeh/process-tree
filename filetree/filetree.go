package filetree

import (
	"fmt"
	"io/fs"
	"path/filepath"

	"github.com/jejikeh/process-tree/tree"
)

type FileTree struct {
	tree.Tree
}

var parentTable = make(map[string]*tree.Node)

func (t *FileTree) AddDir(fullPath string, dir fs.DirEntry) {
	var parent *tree.Node

	parent, foundParent := parentTable[fullPath]

	if !foundParent {
		fmt.Printf("Parent not found for '%s', fallback to root '%s'\n", fullPath, t.Root.Name)

		parent = t.Root
	}

	fullName := fullPath + "/" + dir.Name()

	node := t.Add(fullName, parent)
	parentTable[fullName] = node

	fmt.Printf("Adding node: '%s'\n to parent: '%s'\n", dir.Name(), parent.Name)
}

func (t *FileTree) AddFile(fullPath string, file fs.FileInfo) {
	var parent *tree.Node

	dirPath := filepath.Dir(fullPath)

	parent, foundParent := parentTable[dirPath]

	if !foundParent {
		fmt.Printf("Parent for file not found for '%s', fallback to root '%s'\n", dirPath, t.Root.Name)

		parent = t.Root
	}

	node := t.Add(fullPath, parent)
	node.Size = float64(file.Size())
}

// @Cleanup: This is maybe should go method, once i move the
// Tree to interface and implement FileTree here...
func InitFileTree(dir string) (*FileTree, error) {
	t := &FileTree{}

	// @Note: Root is a special case
	t.initRoot(dir)

	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Skip root node
		if path == dir {
			return nil
		}

		if d.IsDir() {
			parentPath := filepath.Dir(path)
			t.AddDir(parentPath, d)
		} else {
			info, err := d.Info()

			if err != nil {
				return err
			}

			t.AddFile(path, info)
		}

		return nil
	})

	fmt.Printf("Total nodes: '%d'\n", len(t.Nodes))

	for i, value := range t.Nodes {
		if len(value.Childrens) == 0 {
			fmt.Printf("%d	'%s'\n", i, value.Name)
		}
	}

	return t, err
}

func (t *FileTree) initRoot(dir string) {
	dirPath := filepath.Dir(dir)

	fmt.Printf("Adding root node: '%s'", dirPath)

	root := t.Add(dirPath, nil)
	t.Root = root
	parentTable[dirPath] = root
}
