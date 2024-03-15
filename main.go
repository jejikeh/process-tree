package main

import (
	"fmt"

	"github.com/jejikeh/process-tree/filetree"
)

// "fmt"
// "log"
// "os/exec"
// "strings"

// "github.com/jejikeh/process-tree/process"

// gui "github.com/gen2brain/raylib-go/raygui"
// rl "github.com/gen2brain/raylib-go/raylib"

func main() {
	sample := "./samples/file/random/"

	t, err := filetree.InitFileTree(sample)

	if err != nil {
		panic(err)
	}

	fmt.Println()
	fmt.Printf("Total size: \t\t'%f'\n", t.ComputeSize())
	fmt.Printf("Total number of nodes: \t'%d'\n", len(t.Nodes))

	// fmt.Println(t.TreeDump())

	// out, err := exec.Command("ps", "aux").Output()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// lines := strings.Split(string(out), "\n")

	// for i, v := range lines {
	// 	if i == 0 || v == "" {
	// 		continue
	// 	}

	// 	p := process.NewProcess(v)

	// 	if strings.Contains(p.Command, "Code") {
	// 		fmt.Printf("%s", p.ToString())
	// 	}
	// }

	// var button bool

	// w := NewWindow()

	// w.Run(func() {
	// 	button = gui.Button(rl.NewRectangle(50, 150, 100, 40), "Click")
	// 	if button {
	// 		fmt.Println("Clicked on button")
	// 	}
	// })
}
