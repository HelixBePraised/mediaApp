package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

const (
	moviesDirectory = "./media/movies/"
)

type Page struct {
	Title               string
	MovieTitleAndSource map[string]string
}

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob("./temp/*.gohtml"))
}

func main() {
	moviesFileServer := http.FileServer(http.Dir(moviesDirectory))
	http.HandleFunc("/", index)
	http.HandleFunc("/movies/", movieIndex)
	http.Handle("/movies/files/", moviesFileServer)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	p := Page{
		Title: "Home",
	}

	err := tmpl.ExecuteTemplate(w, "header", p)
	check(err, w)

	err = tmpl.ExecuteTemplate(w, "home.gohtml", nil)
	check(err, w)

	err = tmpl.ExecuteTemplate(w, "footer", nil)
	check(err, w)
}

// /movies handler
func movieIndex(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-type", "text/html")

	movMap := getMovieInformation(w)
	p := Page{
		Title:               "Movies",
		MovieTitleAndSource: movMap,
	}

	err := tmpl.ExecuteTemplate(w, "header", p)
	check(err, w)

	err = tmpl.ExecuteTemplate(w, "movieIndex.gohtml", p)
	check(err, w)

	err = tmpl.ExecuteTemplate(w, "footer", nil)
	check(err, w)

}

func check(err error, w http.ResponseWriter) {
	if err != nil {
		fmt.Println(err)
		fmt.Fprintf(w, "%s", err)
	}
}

func getMovieInformation(w http.ResponseWriter) map[string]string {
	m := make(map[string]string)

	err := filepath.Walk(moviesDirectory, func(path string, info os.FileInfo, err error) (er error) {
		if !info.IsDir() {
			src := info.Name()
			name := strings.TrimSuffix(src, filepath.Ext(src))
			fmt.Println(src)
			src = "/movies/files/" + src
			m[name] = src
		}
		return nil
	})

	check(err, w)

	return m
}
