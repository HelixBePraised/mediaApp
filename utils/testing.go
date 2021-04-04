package utils

import (
	"io/fs"
	"path/filepath"
	"time"
)

type FileSystemMock struct {
	FileExistsMock       func(string) bool
	PathIsDirectoryMock  func(string) bool
	GenerateWalkFuncMock func(*[]fs.FileInfo) filepath.WalkFunc
}

func (fs FileSystemMock) FileExists(path string) bool {
	return fs.FileExistsMock(path)
}

func (fs FileSystemMock) PathIsDirectory(path string) bool {
	return fs.PathIsDirectoryMock(path)
}

func (fs FileSystemMock) GenerateWalkFunc(files *[]fs.FileInfo) filepath.WalkFunc {
	return fs.GenerateWalkFuncMock(files)
}

type MockFileInfo struct{}

func (m MockFileInfo) Name() string {
	return "movie.mp4"
}
func (m MockFileInfo) Size() int64 {
	return 1234
}
func (m MockFileInfo) Mode() fs.FileMode {
	return fs.ModeAppend // This is a dummy value for now
}
func (m MockFileInfo) ModTime() time.Time {
	return time.Now()
}
func (m MockFileInfo) IsDir() bool {
	return false
}
func (m MockFileInfo) Sys() interface{} {
	return 1
}
