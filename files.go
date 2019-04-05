package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func GetMediaInformation(path string) map[string]string {
	m := make(map[string]string)

	files, _ := ioutil.ReadDir(path)

	for _, file := range files {
		m[file.Name()] = "/view/" + file.Name()
	}

	// err := filepath.Walk(path, func(walkPath string, info os.FileInfo, err error) error {
	// 	if info.IsDir() {
	// 		name := info.Name()
	// 		m[name] = walkPath
	// 	}
	//
	// 	return nil
	// })

	// if err != nil {
	// 	fmt.Printf("%s", err)
	// }

	return m
}
func GetShowInfo() {
	files, _ := ioutil.ReadDir("./media/shows/")

	for _, dir := range files {

		shows[dir.Name()] = map[string]map[string]string{}

		var seasonName string

		_ = filepath.Walk("./media/shows/"+dir.Name(), func(path string, file os.FileInfo, err error) error {

			if err != nil {
				fmt.Printf("Error: %s\n0", err)
				return nil
			}

			if file.IsDir() {
				seasonName = file.Name()
				shows[dir.Name()][seasonName] = map[string]string{}

			} else {
				title := file.Name()
				title = strings.Replace(title, filepath.Ext(file.Name()), "", -1)
				shows[dir.Name()][seasonName][file.Name()] = path
			}

			return nil
		})
	}
}
