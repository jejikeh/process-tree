package filemap

import (
	"os"
	"path/filepath"
	"testing"
)

const samplePath = "../../samples"

func TestInitTreemapWithCorrectSize(t *testing.T) {
	tree, err := InitTreemap(samplePath)

	if err != nil {
		t.Fatalf("failed to initialized the treemap: %v", err)
	}

	expectedSize, err := dirSize(samplePath)

	if err != nil {
		t.Fatalf("failed to compute size of dir: %v", err)
	}

	if tree.ComputeSizes() != expectedSize {
		t.Fatalf("expected size is %f, but got %f", expectedSize, tree.ComputeSizes())
	}
}

func TestInitTreemapWithCorrectLength(t *testing.T) {
	tree, err := InitTreemap(samplePath)

	if err != nil {
		t.Fatalf("failed to initialized the treemap: %v", err)
	}

	expectedSize, err := dirFilesCount(samplePath)

	if err != nil {
		t.Fatalf("failed to compute count of files: %v", err)
	}

	if len(tree.Nodes) != expectedSize {
		t.Fatalf("expected count of files  is %d, but got %d", expectedSize, len(tree.Nodes))
	}
}

func dirSize(path string) (float64, error) {
	var size float64

	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			size += float64(info.Size())
		}

		return err
	})
	return size, err
}

func dirFilesCount(path string) (int, error) {
	var size int

	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		size++

		return err
	})

	return size, err
}
