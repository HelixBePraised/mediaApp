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
	mediaDirectory  = "./media/"
	moviesDirectory = "./media/movies/"
	showsDirectory  = "./media/shows/"
)

type Page struct {
	Title               string
	MediaTitleAndSource map[string]string
	MediaSrc            string
}

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob("./temp/*.gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/movies/", movieIndex)
	http.Handle("/files/", http.StripPrefix("/files/", http.FileServer(http.Dir(mediaDirectory))))
	http.HandleFunc("/viewmovie/", movieViewerHandler)
	http.HandleFunc("/shows", showsIndex)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	p := Page{
		Title: "Home",
	}

	err := tmpl.ExecuteTemplate(w, "home.gohtml", p)
	check(err, w)

}

// /movies handler
func movieIndex(w http.ResponseWriter, req *http.Request) {

	movMap := getMediaInformation(w, moviesDirectory)

	p := Page{
		Title:               "Movies",
		MediaTitleAndSource: movMap,
	}

	err := tmpl.ExecuteTemplate(w, "movieIndex.gohtml", p)
	check(err, w)

}

func showsIndex(w http.ResponseWriter, req *http.Request) {

	seasonInfo := getMediaInformation(w, showsDirectory)

	p := Page{
		Title:               "Shows",
		MediaTitleAndSource: seasonInfo,
	}

	err := tmpl.ExecuteTemplate(w, "movieIndex.gohtml", p)
	check(err, w)
}

func movieViewerHandler(w http.ResponseWriter, req *http.Request) {
	url := req.URL.Path

	var title string
	title = strings.TrimSuffix(url, filepath.Ext(url))
	title = strings.Replace(title, "/movies/", "/files/", -1)
	title = strings.Replace(title, filepath.Ext(title), "", -1)
	title = strings.Replace(title, "/viewmovie/", "", -1)
	url = strings.Replace(url, "/viewmovie/", "/files/", -1)

	url = "http://localhost:8080" + url

	p := Page{
		Title:    title,
		MediaSrc: url,
	}

	err := tmpl.ExecuteTemplate(w, "viewer.gohtml", p)
	check(err, w)

}

//Fine for now, possibly want to do more with errors
func check(err error, w http.ResponseWriter) {
	if err != nil {
		fmt.Println(err)
		fmt.Fprintf(w, "%s", err)
	}
}

func getMediaInformation(w http.ResponseWriter, path string) map[string]string {
	m := make(map[string]string)

	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) (er error) {
		name := info.Name()
		m[name] = path + name
		return nil
	})

	check(err, w)

	return m
}
