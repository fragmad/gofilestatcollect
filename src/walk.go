package main

import "path/filepath"
import "os"
import "io/ioutil"
import "fmt"
import "time"

//import "gopkg.in/mgo.v2"

var count = 0
var file_count = 0
var directory_count = 0

var verbose = false
var mongo = false
var start_time = time.Now()

// Variable to store list of paths to ignore and not walk.
// Ensure /proc is in this list.

var ignored_paths = [1]string{"/proc"}

func isPathinIgnoreList(path string) bool {
	ignore_path := false
	for _, test_path := range ignored_paths {
		if path == test_path {
			fmt.Println("ignoring", test_path)
			ignore_path = true
		}
	}
	return ignore_path
}

/* type Entry struct {
	FullPath    string  `bson:"FullPath"`
	Name        string  `bson:"Name"`
	Size        string  `bson:"Size"`
	IsDirectory boolean `bson:"IsDirectory"`
	IsDirectory boolean `bson:"IsFile"`
}*/

func print_file_stats(path string, info os.FileInfo) {
	fmt.Println("--")
	fmt.Print(path)
	fmt.Println(info.Name())
	fmt.Printf("%d Bytes. \n", info.Size())
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

func print_final_stats() {
	fmt.Printf("%d items found.\n", count)
	fmt.Printf("%d files found.\n", file_count)
	fmt.Printf("%d directories found.\n", directory_count)
	fmt.Printf("Execution took: %s.\n", time.Since(start_time))
}

func walk_file_tree(dirPath string) {
	filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}

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
	path := "."
	if len(os.Args) > 1 {
		path = os.Args[1]
	}

	walk_file_tree(path)
	print_final_stats()
}
