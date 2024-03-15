package tree

import (
	"fmt"
	"io/fs"
	"log"
	"path/filepath"
	"strings"
)

type DisplayTree struct {
	Tree

	DisplayNodes []DisplayNode
}

type Tree struct {
	Nodes []*Node
	Root  *Node
}

func (t *Tree) Add(name string, parent *Node) *Node {
	// @Cleanup: Replace all panics with errors
	if len(t.Nodes) == 0 {
		if parent != nil {
			log.Fatalf("%s root node cannot have a parent", name)
		}
	} else {
		if parent == nil {
			log.Fatalf("%s node must have a parent", name)
		}
	}

	node := &Node{}
	node.Name = name
	node.Parent = parent

	t.Nodes = append(t.Nodes, node)

	if parent != nil {
		parent.Childrens = append(parent.Childrens, node)
	}

	return node
}

var parentTable = make(map[string]*Node)

// @Cleanup: Maybe Tree should be interface and then theese will go
// to FileTree, for process logic it will be ProcessTree
func (t *Tree) AddDir(fullPath string, name string) {
	var parent *Node

	parent, foundParent := parentTable[fullPath]

	if !foundParent {
		fmt.Printf("Parent not found for '%s', fallback to root '%s'\n", name, t.Root.Name)

		parent = t.Root
	}

	fullName := fullPath + "/" + name

	node := t.Add(fullName, parent)
	parentTable[fullName] = node

	fmt.Printf("Adding node: '%s'\n to parent: '%s'\n", name, parent.Name)
}

func (t *Tree) AddFile(fullname string) {

}

// @Cleanup: This is maybe should go method, once i move the
// Tree to interface and implement FileTree here...
func InitTree(dir string) (*Tree, error) {
	t := &Tree{}

	// @Note: Root is a special case
	t.initRoot(dir)

	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			parentPath := filepath.Dir(path)

			t.AddDir(parentPath, d.Name())
		} else {
			t.AddFile(path)
		}

		return nil
	})

	fmt.Printf("Total nodes: '%d'", len(t.Nodes))

	return t, err
}

func (t *Tree) initRoot(dir string) {
	dirPath := filepath.Dir(dir)

	fmt.Printf("Adding root node: '%s'", dirPath)

	root := t.Add(dirPath, nil)
	t.Root = root
	parentTable[dirPath] = root
}

// @Incomplete: This is not implemented.
func (t *Tree) TreeDump() string {
	sb := strings.Builder{}

	for _, node := range t.Nodes {
		sb.WriteString(node.Name)
		sb.WriteRune('\n')
	}

	return sb.String()
}
