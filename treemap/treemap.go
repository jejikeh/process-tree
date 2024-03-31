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

type TreemapNode struct {
	Name      string
	Size      float64
	Parent    *TreemapNode
	Childrens []*TreemapNode
}

type Treemap struct {
	Nodes []TreemapNode
}

func NewTreemap() Treemap {
	return Treemap{
		Nodes: make([]TreemapNode, 0),
	}
}

func (t *Treemap) Add(name string, parent *TreemapNode) (*TreemapNode, error) {
	if len(t.Nodes) == 0 && parent != nil {
		return nil, fmt.Errorf("the root node cannot have a parent")
	}

	if len(t.Nodes) != 0 && parent == nil {
		return nil, fmt.Errorf("the non-root node must have a parent")
	}

	node := TreemapNode{
		Name:   name,
		Size:   0,
		Parent: parent,
	}

	t.Nodes = append(t.Nodes, node)

	if parent != nil {
		// @Incomplete: Maybe i can use just node here?
		parent.Childrens = append(parent.Childrens, &t.Nodes[len(t.Nodes)-1])
	}

	return &t.Nodes[len(t.Nodes)-1], nil
}
