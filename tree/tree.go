package tree

import (
	"log"
	"sort"
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type DisplayTree struct {
	Tree

	DisplayNodes []DisplayNode
	Dirty        bool
}

func InitDisplayTree(tree *Tree) *DisplayTree {
	return &DisplayTree{
		Tree: *tree,
	}
}

func (t *DisplayTree) ComputeSize(size rl.Vector2) {
	if !t.Dirty {
		return
	}

	t.Dirty = false

	if len(t.Nodes) == 0 || t.Root.Size == 0.0 {
		return
	}

	t.DisplayNodes = make([]DisplayNode, 0, len(t.Nodes))

	rootSize := t.Root.Size

	avaiable := size

	totalArea := float64(size.X * size.Y)

	for i, node := range t.Nodes {
		fraction := node.Size / rootSize
		area := totalArea * fraction

		width := area / float64(size.Y)
		height := size.Y

		d := t.DisplayNodes[i]

		d.Corner = rl.Vector2{X: float32(avaiable.X), Y: float32(avaiable.Y)}
		d.Size = rl.Vector2{X: float32(width), Y: float32(height)}

		avaiable.X -= float32(width)
		avaiable.Y -= float32(height)
	}

}

type Tree struct {
	Nodes []*Node
	Root  *Node

	Dirty bool
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
	if !t.Dirty {
		return 0.0
	}

	t.Dirty = false

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

	// Pass 3: Sort by size
	for i := len(t.Nodes) - 1; i >= 0; i-- {
		node := t.Nodes[i]

		if len(node.Childrens) != 0 {
			sort.Slice(node.Childrens, func(i, j int) bool {
				return node.Childrens[i].Size > node.Childrens[j].Size
			})
		}
	}

	return t.Root.Size
}
