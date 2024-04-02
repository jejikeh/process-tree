package treemap

import (
	"fmt"
	"testing"
)

func TestAddWrongNodes(t *testing.T) {
	tree := NewTreemap()

	// cant add root-node with parent
	rootNode := "root_node"
	_, err := tree.Add(rootNode, &Node{})

	if err == nil {
		t.Fatalf("root node cant have a parent")
	}

	// should add node
	root, err := tree.Add(rootNode, nil)

	if err != nil {
		t.Fatalf("error while adding node: %v", err)
	}

	if root == nil {
		t.Fatalf("the root node is not initialized")
	}

	// cant add non-root node without parent
	testNodeName := "test_node"

	_, err = tree.Add(testNodeName, nil)

	if err == nil {
		t.Fatalf("non-root node must have a parent")
	}
}

func TestAddNode(t *testing.T) {
	tree := NewTreemap()

	testNode := "test_node"

	node, err := tree.Add(testNode, nil)

	if err != nil {
		t.Fatalf("error adding node: %v", err)
	}

	if len(tree.Nodes) != 1 {
		t.Fatal("node was not added to the treemap")
	}

	if node.Name != testNode {
		t.Fatalf("the added node name need to be %s, but got %s", testNode, node.Name)
	}

	if node != tree.Root {
		t.Fatalf("the pointer to node returned from tree.Add() should be the same as node in tree Nodes field")
	}
}

func TestComputeSizesFromRootToRoot(t *testing.T) {
	nodeCount := 10
	nodeSize := 100.0

	tree := NewTreemap()

	if err := seedTreeFromRootToRoot(&tree, nodeCount, nodeSize); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expectedSize := float64(nodeCount) * nodeSize

	if tree.ComputeSizes() != expectedSize {
		t.Fatalf("the expected size was %f, but got %f", expectedSize, tree.ComputeSizes())
	}
}

func TestComputeSizesFromRootToNodes(t *testing.T) {
	nodeCount := 10
	nodeSize := 100.0

	tree := NewTreemap()

	seedTreeFromRootToNodesWhenIndexIsEven(&tree, nodeCount, nodeSize)

	expectedSize := float64(nodeCount) * nodeSize * 5

	if tree.ComputeSizes() != expectedSize {
		t.Fatalf("the expected size was %f, but got %f", expectedSize, tree.ComputeSizes())
	}
}

func TestReComputeSizes(t *testing.T) {
	nodeCount := 10
	nodeSize := 100.0
	nodesToDelete := 3

	tree := NewTreemap()

	seedTreeFromRootToNodesWhenIndexIsEven(&tree, nodeCount, nodeSize)

	expectedSize := float64(nodeCount) * nodeSize * 5

	if tree.ComputeSizes() != expectedSize {
		t.Fatalf("the expected size was %f, but got %f", expectedSize, tree.ComputeSizes())
	}

	tree.Nodes = tree.Nodes[:len(tree.Nodes)-nodesToDelete-1]
	expectedSizeAfterDelete := expectedSize - float64(nodesToDelete)*nodeSize

	if tree.ComputeSizes() != expectedSize {
		t.Fatalf("the expected size was %f, but got %f", expectedSize, tree.ComputeSizes())
	}

	tree.ReComputeSizes()

	if tree.ComputeSizes() != expectedSizeAfterDelete {
		t.Fatalf("the expected size was %f, but got %f", expectedSize, tree.ComputeSizes())
	}
}

func seedTreeFromRootToRoot(tree *Treemap, nodeCount int, nodeSize float64) error {
	rootNode, err := tree.Add("root", nil)

	if err != nil {
		return fmt.Errorf("got unexpected error: %v", err)
	}

	for i := 0; i < nodeCount; i++ {
		node, err := tree.Add(fmt.Sprintf("node_%d", i), rootNode)

		if err != nil {
			return fmt.Errorf("got unexpected error: %v", err)
		}

		node.Size = nodeSize
	}

	return nil
}

func seedTreeFromRootToNodesWhenIndexIsEven(tree *Treemap, nodeCount int, nodeSize float64) error {
	rootNode, err := tree.Add("root", nil)

	if err != nil {
		return err
	}

	node := rootNode

	for i := 0; i < nodeCount; i++ {
		if i%2 == 0 {
			if err = seedNode(tree, node, nodeCount, nodeSize); err != nil {
				return err
			}
		} else {
			node, err = tree.Add(fmt.Sprintf("node_%d", i), rootNode)

			if err != nil {
				return err
			}
		}
	}

	return nil
}

func seedNode(tree *Treemap, parentNode *Node, childCount int, nodeSize float64) error {
	for i := 0; i < childCount; i++ {
		node, err := tree.Add(fmt.Sprintf("%s_%d", parentNode.Name, i), parentNode)

		if err != nil {
			return fmt.Errorf("got unexpected error: %v", err)
		}

		node.Size = nodeSize
	}

	return nil
}
