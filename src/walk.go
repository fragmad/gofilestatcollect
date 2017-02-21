package main

import "path/filepath"
import "os"
import "io/ioutil"
import "fmt"

var count = 0
var file_count = 0
var directory_count = 0
var verbose = false

func print_file_stats(path string, info os.FileInfo) {
	fmt.Println("--")
	fmt.Print(path)
	fmt.Println(info.Name())
	fmt.Print(info.Size())
	fmt.Println(" Bytes")
	fmt.Println("Not a directory")
}

func print_directory_stats(path string, info os.FileInfo) {
	fmt.Println("--")
	fmt.Print(path)
	fmt.Println(info.Name())
	files, _ := ioutil.ReadDir(path)
	fmt.Print(len(files))
	fmt.Println(" files")
	fmt.Println("A directory")
}

func walk_file_tree(dirPath string) {
	filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			if verbose {
				print_file_stats(path, info)
			}
			count++
			file_count++
		}
		if info.IsDir() {
			if verbose {
				print_directory_stats(path, info)
			}
			count++
			directory_count++
		}
		return nil
	})
}

func main() {
	path := os.Args[1]

	walk_file_tree(path)
	fmt.Printf("%d items found.\n", count)
	fmt.Printf("%d files found.\n", file_count)
	fmt.Printf("%d directories found.\n", directory_count)
}
