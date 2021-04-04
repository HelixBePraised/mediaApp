package main

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"reflect"
	"testing"

	"git.jacksontaylor.xyz/jama/utils"
)

func TestReadDirectory(t *testing.T) {
	// Mock where fs does not find a file
	pathDoesNotExist := func(_ string) bool {
		return false
	}

	// Mock where fs finds the path
	pathDoesExist := func(_ string) bool {
		return true
	}

	// Mock where the path is a file and not a directory
	pathIsNotDirectory := func(_ string) bool {
		return false
	}

	// Mock where the path is a directory
	pathIsDirectory := func(_ string) bool {
		return true
	}

	// Mock where the walkFunc returns an error for some reason
	walkFuncReturnError := func(_ *[]fs.FileInfo) filepath.WalkFunc {
		return func(_ string, _ fs.FileInfo, _ error) error {
			return fmt.Errorf("Some ominous file system error!")
		}
	}

	walkFuncReturnsFiles := func(files *[]fs.FileInfo) filepath.WalkFunc {
		*files = []fs.FileInfo{utils.MockFileInfo{}}
		return func(_ string, _ fs.FileInfo, _ error) error {
			return nil
		}
	}

	type args struct {
		fsi  FileSystemInteractor
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    []fs.FileInfo
		wantErr bool
	}{
		{
			name: "fs is nil",
			args: args{
				fsi:  nil,
				path: "/somePath",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Missing Path",
			args: args{
				fsi:  utils.FileSystemMock{FileExistsMock: pathDoesNotExist},
				path: "/incorrectPath",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Path is not a directory",
			args: args{
				fsi: utils.FileSystemMock{
					FileExistsMock:      pathDoesExist,
					PathIsDirectoryMock: pathIsNotDirectory,
				},
				path: "/aFile.txt",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Walkfunc returns error",
			args: args{
				fsi: utils.FileSystemMock{
					FileExistsMock:       pathDoesExist,
					PathIsDirectoryMock:  pathIsDirectory,
					GenerateWalkFuncMock: walkFuncReturnError,
				},
				path: "/somePath",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "WalkFunc returns files",
			args: args{
				fsi: utils.FileSystemMock{
					FileExistsMock:       pathDoesExist,
					PathIsDirectoryMock:  pathIsDirectory,
					GenerateWalkFuncMock: walkFuncReturnsFiles,
				},
				path: "/somePath",
			},
			want:    []fs.FileInfo{utils.MockFileInfo{}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadDirectory(tt.args.fsi, tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadDirectory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadDirectory() = %v, want %v", got, tt.want)
			}
		})
	}
}
