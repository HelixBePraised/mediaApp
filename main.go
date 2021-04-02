package main

import (
	"fmt"
	"os"
	// "path/filepath"
	"io/fs"
)

func main() {
	fmt.Println("Hello world!")
}

type FileSystemInteractor interface {
	FileExists(string) bool
	PathIsDirectory(string) bool
}

type FileSystem struct{}

func (f FileSystem) FileExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

func (f FileSystem) PathIsDirectory(path string) bool {
	stat, err := os.Stat(path)
	if err != nil || !stat.IsDir() {
		return false
	}
	return true
}

// Read a directory and return all the filesnames in it
func ReadDirectory(fsi FileSystemInteractor, path string) ([]fs.FileInfo, error) {
	if fsi == nil || !fsi.FileExists(path) || !fsi.PathIsDirectory(path) {
		return nil, fmt.Errorf("%s is not a valid path!", path)
	}

	// Get the contents of the directory
	return []fs.FileInfo{}, nil
}
