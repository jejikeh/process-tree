package tree

import (
	"log"
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

	// @Note: If there are Parent, so parent is a directory
	if parent != nil {
		parent.Childrens = append(parent.Childrens, node)
	}

	return node
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

func (t *Tree) ComputeSize() float64 {
	// Pass 1: Cleanup sizes for dirs.
	for _, node := range t.Nodes {
		// If node has children, then it is a dir
		if len(node.Childrens) != 0 {
			node.Size = 0
		}
	}

	// Pass 2: Iterate from leaves to root
	for i := len(t.Nodes) - 1; i >= 0; i-- {
		node := t.Nodes[i]

		if node.Parent != nil {
			node.Parent.Size += node.Size
		}
	}

	return t.Root.Size
}
