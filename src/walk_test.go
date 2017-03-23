package main

import (
	"os"
	"testing"
)

func TestIsPathInIgnoreList(t *testing.T) {
	result := isPathinIgnoreList("/proc")
	expected := true

	if result != expected {
		t.Error(result, expected)
	}

}

func TestWalkFileTree(t *testing.T) {
	os.MkdirAll("./test_dirs/test1", os.ModePerm)
	os.Create("./test_dirs/test1.txt")
	os.Create("./test_dirs/test1/test2.txt")

	result := walk_file_tree("./test_dirs")
	expected := 4

	if result != expected {
		t.Error(result, expected)
	}

	os.RemoveAll("./test_dirs/")
}
