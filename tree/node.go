package tree

import rl "github.com/gen2brain/raylib-go/raylib"

type DisplayNode struct {
	Index int
	Color rl.Color

	Corner rl.Vector2
	Size   rl.Vector2
}

type Node struct {
	Name string
	Size float64

	Parent    *Node
	Childrens []*Node

	Index int
}
