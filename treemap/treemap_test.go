package treemap

import "testing"

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
