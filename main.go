package main

import (
	"fmt"

	"github.com/jejikeh/process-tree/treemap/filemap"
)

func main() {
	tree, err := filemap.InitTreemap("samples")

	if err != nil {
		panic(err)
	}

	for _, node := range tree.Nodes {
		for _, file := range node.Childrens {
			fmt.Printf("FILE \n\t%s \n\t[%f]\n", file.Name, file.Size)
		}
	}

	fmt.Printf("\n\nFile count: [%d]\n", len(tree.Nodes))
	fmt.Printf("\n\nCalculated size: [%f]\n", tree.ComputeSizes())
}
