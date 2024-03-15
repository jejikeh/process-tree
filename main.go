package main

import (
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

	window.Font = rl.LoadFont("assets/fonts/Martel-Regular.ttf")

	w.EntityManager.Add(entity.NewTreemap(&t.Tree))

	w.Run(func() {
		w.EntityManager.Update()
	}, func() {
		rl.ClearBackground(rl.RayWhite)

		w.EntityManager.Draw()
	})
}
