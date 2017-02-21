package main

import "path/filepath"
import "os"

import "io/ioutil"

func walk_file_tree(dirPath string) {
	filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			println("--")
			print(path)
			println(info.Name())
			print(info.Size())
			println(" Bytes")
			println("Not a directory")
		}
		if info.IsDir() {
			println("--")
			print(path)
			println(info.Name())
			files, _ := ioutil.ReadDir(path)
			print(len(files))
			println(" files")
			println("A directory")
		}
		return nil
	})
}

func main() {
	path := os.Args[1]

	walk_file_tree(path)
}
