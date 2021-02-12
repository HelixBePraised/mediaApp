package main

import (
	"fmt"
	"os"
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

// Read a directory and return all the filesnames in it
func ReadDirectory(fs FileSystemInteractor, path string) ([]string, error) {
	if fs == nil || !fs.FileExists(path) {
		return nil, fmt.Errorf("%s is not a valid path!", path)
	}
	return []string{}, nil
}
