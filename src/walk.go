package main

import "path/filepath"
import "os"
import "io/ioutil"

const dirPath = "/home/lol"

func main() {
	// walk all files in directory
	filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			println("--")
			println(info.Name())
			print(info.Size())
			println(" Bytes")
			println("Not a directory")
		}
		if info.IsDir() {
			println("--")
			println(info.Name())
			/*		files, _ := ioutil.ReadDir(path)
					print(len(files))
					println(" files")*/
			println("A directory")
		}
		return nil
	})
}
