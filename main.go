package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/jejikeh/process-tree/entity"
	"github.com/jejikeh/process-tree/filetree"
	"github.com/jejikeh/process-tree/window"
)

func main() {
	sample := "./samples/file/random/"

	t, err := filetree.InitFileTree(sample)

	if err != nil {
		panic(err)
	}

	w := window.NewWindow()

	entity.Font = rl.LoadFont("assets/fonts/Martel-Regular.ttf")

	l := entity.NewText(rl.Vector2{X: 12, Y: 436}, fmt.Sprintf("Total size: \t\t'%f'\nTotal number of nodes: \t'%d'", t.ComputeSize(), len(t.Nodes)))

	w.EntityManager.Add(l)

	w.EntityManager.Add(entity.NewTreemap(&t.Tree))

	w.Run(func() {
		w.EntityManager.Update()
	}, func() {
		rl.ClearBackground(rl.RayWhite)

		w.EntityManager.Draw()
	})
}
