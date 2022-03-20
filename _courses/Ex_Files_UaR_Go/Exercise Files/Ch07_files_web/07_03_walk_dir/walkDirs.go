package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {

	root, err := filepath.Abs(".")
	checkError(err)
	fmt.Println("Working path: ", root)

	err = filepath.Walk(root, processPath)
	checkError(err)
}

func checkError(err error){
	if err != nil {
		panic(err)
	}
}

func processPath(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	if path != "." {
		if info.IsDir() {
			fmt.Println("Directory: ", path)
		} else {
			fmt.Println("File:", path)
		}
	}

	return nil
}