package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func displayFolderStructure(path string, indent string) {
	if filepath.Base(path) == ".git" {
		return
	}

	fmt.Println(indent+"|--", filepath.Base(path))

	fileInfo, err := os.Stat(path)
	if err != nil || !fileInfo.IsDir() {
		return
	}

	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening directory:", err)
		return
	}
	defer file.Close()

	files, err := file.Readdir(0)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	for _, fileInfo := range files {
		childPath := filepath.Join(path, fileInfo.Name())
		displayFolderStructure(childPath, indent+"   ")
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run folder_structure.go <path>")
		return
	}

	path := os.Args[1]
	fmt.Println("Folder Structure for:", path)
	displayFolderStructure(path, "")
}
