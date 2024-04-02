// package main

// import (
// 	// "fmt"

// 	rl "github.com/gen2brain/raylib-go/raylib"
// 	// "github.com/jejikeh/process-tree/treemap/filemap"
// )

// func main() {
// 	// tree, err := filemap.InitTreemap("samples")

// 	// if err != nil {
// 	// 	panic(err)
// 	// }

// 	// for _, node := range tree.Nodes {
// 	// 	for _, file := range node.Childrens {
// 	// 		fmt.Printf("FILE \n\t%s \n\t[%f]\n", file.Name, file.Size)
// 	// 	}
// 	// }

// 	// fmt.Printf("\n\nFile count: [%d]\n", len(tree.Nodes))
// 	// fmt.Printf("\n\nCalculated size: [%f]\n", tree.ComputeSizes())

// 	// tree.ReComputeSizes()

// 	rl.InitWindow(800, 450, "raylib [core] example - basic window")
// 	defer rl.CloseWindow()

// 	rl.SetTargetFPS(60)

// 	for !rl.WindowShouldClose() {
// 		rl.BeginDrawing()

// 		rl.ClearBackground(rl.RayWhite)
// 		rl.DrawText("Congrats! You created your first window!", 190, 200, 20, rl.LightGray)

// 		rl.EndDrawing()
// 	}
// }

package main

import (
	"fmt"

	gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(800, 450, "raygui - button")

	rl.SetTargetFPS(60)

	var button bool

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)

		button = gui.Button(rl.NewRectangle(50, 150, 100, 40), "Click")
		if button {
			fmt.Println("Clicked on button")
		}

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
