package treemap

import (
	"fmt"
	"image/color"
)

type DisplayNode struct {
	Color color.RGBA
}

type TreemapDisplay struct {
}

type Node struct {
	Name      string
	Size      float64
	Parent    *Node
	Childrens []*Node
}

type Treemap struct {
	Nodes []*Node

	Root *Node
}

func NewTreemap() Treemap {
	return Treemap{
		Nodes: make([]*Node, 0),
	}
}

func (t *Treemap) Add(name string, parent *Node) (*Node, error) {
	if len(t.Nodes) == 0 && parent != nil {
		return nil, fmt.Errorf("the root node cannot have a parent")
	}

	if len(t.Nodes) != 0 && parent == nil {
		return nil, fmt.Errorf("the non-root node must have a parent")
	}

	node := &Node{
		Name:   name,
		Size:   0,
		Parent: parent,
	}

	t.Nodes = append(t.Nodes, node)

	if parent != nil {
		parent.Childrens = append(parent.Childrens, t.Nodes[len(t.Nodes)-1])
	}

	if len(t.Nodes) == 1 {
		t.Root = t.Nodes[len(t.Nodes)-1]
	}

	return t.Nodes[len(t.Nodes)-1], nil
}

func (t *Treemap) ComputeSizes() float64 {
	if t.Root.Size == 0.0 {
		return t.ReComputeSizes()
	}

	return t.Root.Size
}

func (t *Treemap) ReComputeSizes() float64 {
	for _, node := range t.Nodes {
		// @Note: This is directory
		if len(node.Childrens) != 0 {
			node.Size = 0
		}
	}

	for i := len(t.Nodes) - 1; i >= 1; i-- {
		node := t.Nodes[i]

		fmt.Println(node.Parent.Name)

		node.Parent.Size += node.Size
	}

	return t.Root.Size
}
