package main

import (
	"reflect"
	"testing"
)

type fileSystemMock struct {
	fileExistsMock func(string) bool
}

func (fs fileSystemMock) FileExists(path string) bool {
	return fs.fileExistsMock(path)
}

func TestReadDirectory(t *testing.T) {
	// Mock where fs does not find a file
	pathDoesNotExist := func(_ string) bool {
		return false
	}

	type args struct {
		fs   FileSystemInteractor
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "Missing Path",
			args: args{
				fs:   fileSystemMock{fileExistsMock: pathDoesNotExist},
				path: "/incorrectPath",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadDirectory(tt.args.fs, tt.args.path)
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
