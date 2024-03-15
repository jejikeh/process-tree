package entity

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/jejikeh/process-tree/tree"
)

type Treemap struct {
	*Transform

	RenderRect  rl.Rectangle
	Tree        *tree.Tree
	DisplayTree *tree.DisplayTree
}

func NewTreemap(t *tree.Tree) *Treemap {
	tr := &Treemap{
		Transform: &Transform{
			Position: rl.Vector2{X: 12, Y: 12},
			Size:     rl.Vector2{X: 640 - 24, Y: 480 - 64},
		},
		Tree:        t,
		DisplayTree: tree.InitDisplayTree(t),
	}

	tr.RenderRect = rl.NewRectangle(tr.Position.X, tr.Position.Y, tr.Size.X, tr.Size.Y)

	return tr
}

func (t *Treemap) Draw() {
	t.Tree.ComputeSize()
	t.DisplayTree.ComputeSize(t.Size)

	rl.DrawRectangleGradientEx(t.RenderRect, rl.Red, rl.Orange, rl.Black, rl.Orange)
}

func (t *Treemap) Update() {
}
