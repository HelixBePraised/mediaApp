package main

import (
	_ "reflect"
	"testing"
)

type fileSystemMock struct{}

var fileExistsMock func(string) bool

func (fs fileSystemMock) FileExists(path string) bool {
	return fileExistsMock(path)
}

func TestReadDirectory(t *testing.T) {
	testPath := "/incorrectPath/"

	fs := fileSystemMock{}

	// FileSystem can't find the file
	fileExistsMock = func(_ string) bool {
		return false
	}

	_, err := ReadDirectory(fs, testPath)

	if err == nil {
		t.Error("Error was not returned!")
	}
}
