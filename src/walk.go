package main

import "path/filepath"
import "os"
import "io/ioutil"
import "fmt"

//import "gopkg.in/mgo.v2"

var count = 0
var file_count = 0
var directory_count = 0
var verbose = true
var mongo = false

type Entry struct {
	FullPath    string  `bson:"FullPath"`
	Name        string  `bson:"Name"`
	Size        string  `bson:"Size"`
	IsDirectory boolean `bson:"IsDirectory"`
	IsDirectory boolean `bson:"IsFile"`
}

var mongo boolean

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

// Do the connection stuff for mongo

// All disgustingly ineffiecent and hardcoded.

func insert_into_mongod(path_item Entry) {

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
				sssssssss
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
