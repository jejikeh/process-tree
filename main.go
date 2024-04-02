// package main
// import (

// 	// "fmt"

// 	rl "github.com/gen2brain/raylib-go/raylib"
// 	// "github.com/jejikeh/process-tree/treemap/filemap"
// )

// func main() {
// 	//

// 	// fmt.Printf("\n\nFile count: [%d]\n", len(tree.Nodes))
// 	// fmt.Printf("\n\nCalculated size: [%f]\n", tree.ComputeSizes())

// 	// tree.ReComputeSizes()

// 	rl.InitWindow(800, 450, "raylib [core] example - basic window")
// 	defer rl.CloseWindow()

// 	rl.SetTargetFPS(60)

// 	for !rl.WindowShouldClose() {
// 		rl.BeginDrawing()

// 		rl.ClearBackground(rl.RayWhite)

// 		rl.EndDrawing()
// 	}
// }

package main

import (
	"fmt"

	gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/jejikeh/process-tree/treemap/filemap"
)

func main() {
	rl.InitWindow(800, 450, "raygui - button")

	rl.SetTargetFPS(60)

	tree, err := filemap.InitTreemap("samples")

	if err != nil {
		panic(err)
	}

	button := false
	var scrollY float32 = 0.0

	for !rl.WindowShouldClose() {
		scrollY += rl.GetMouseWheelMove()

		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)

		// rl.DrawText(fmt.Sprintf("Total Size: %f", tree.ComputeSizes()), 190, 200, 20, rl.LightGray)

		for i, node := range tree.Nodes {
			rl.DrawText(fmt.Sprintf("[%s] =%f", node.Name, node.Size), 10, int32(39+((i+1)*20)+int(scrollY)), 20, rl.LightGray)
		}

		button = gui.Button(rl.NewRectangle(10, float32(10+(int(scrollY))), 100, 40), "Compute Sizes")
		if button {
			tree.ComputeSizes()
		}

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
