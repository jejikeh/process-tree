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
	return &Treemap{
		Transform: &Transform{
			Position: rl.Vector2{X: float32(rl.GetScreenWidth()) / 2, Y: float32(rl.GetScreenHeight()) / 2},
			Size:     rl.Vector2{X: 640 - 24, Y: 480 - 64},
		},
		RenderRect:  rl.NewRectangle(12, 12, float32(rl.GetScreenWidth())-24, float32(rl.GetScreenHeight())-64),
		Tree:        t,
		DisplayTree: tree.InitDisplayTree(t),
	}
}

func (t *Treemap) Draw() {
	t.DisplayTree.ComputeSize(t.Size)
	t.Tree.ComputeSize()

	rl.DrawRectangleGradientEx(t.RenderRect, rl.Red, rl.Orange, rl.Black, rl.Orange)
}

func (t *Treemap) Update() {
}
