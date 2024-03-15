package tree

import rl "github.com/gen2brain/raylib-go/raylib"

type DisplayNode struct {
	Index int
	Color rl.Color
}

type Node struct {
	Name string
	Size float32

	Parent    *Node
	Childrens []*Node

	Index int
}
