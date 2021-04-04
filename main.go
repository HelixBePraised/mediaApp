package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func main() {
	fmt.Println("Hello world!")
}

type FileSystemInteractor interface {
	FileExists(string) bool
	PathIsDirectory(string) bool
	GenerateWalkFunc(*[]fs.FileInfo) filepath.WalkFunc
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

// GenerateWalkFunc generates a walk func that appends each file found to a slice of files
func (f FileSystem) GenerateWalkFunc(files *[]fs.FileInfo) filepath.WalkFunc {
	return func(_ string, info fs.FileInfo, _ error) error {
		if !info.IsDir() {
			*files = append(*files, info)
		}
		return nil
	}
}

// Read a directory and return all the filesnames in it
func ReadDirectory(fsi FileSystemInteractor, path string) ([]fs.FileInfo, error) {
	if fsi == nil {
		return nil, fmt.Errorf("FileSystemInteractor is nil!")
	}

	if !fsi.FileExists(path) || !fsi.PathIsDirectory(path) {
		return nil, fmt.Errorf("%s is not a valid path!", path)
	}

	var files []fs.FileInfo

	err := filepath.Walk(path, fsi.GenerateWalkFunc(&files))

	if err != nil {
		return nil, fmt.Errorf("Error occurred in walk func: %s", path)
	}

	// Get the contents of the directory
	return files, nil
}
