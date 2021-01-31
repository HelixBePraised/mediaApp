package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	mediaDirectory  = "./media"
	moviesDirectory = "./media/movies/."
	//showsDirectory  = "./media/shows/"
)

var movies = map[string]string{}
var shows = map[string]map[string]map[string]string{}

type Page struct {
	Title             string
	MediaTitleAndLink map[string]string
	MediaSrc          string
}

var tmpl *template.Template

func init() {
	fmt.Print("[+] Parsing templates...")
	tmpl = template.Must(template.ParseGlob("./temp/*.gohtml"))
	fmt.Print(" Done\n")
	fmt.Print("[+] Gathering Movie Map...")
	movies = GetMediaInformation(moviesDirectory)
	fmt.Print(" Done\n")
	fmt.Print("[+] Gathering Show Map...")
	GetShowInfo()
	fmt.Print(" Done\n")
}

func main() {
	fmt.Println("[+] Serving at localhost:8080")
	r := mux.NewRouter()

	r.HandleFunc("/shows", showsHandler)
	r.HandleFunc("/shows/", showsHandler)
	r.HandleFunc("/shows/{show}", showsHandler)
	r.HandleFunc("/shows/{show}/", showsHandler)
	r.HandleFunc("/shows/{show}/{season}", showsHandler)
	r.HandleFunc("/shows/{show}/{season}/", showsHandler)
	r.HandleFunc("/shows/{show}/{season}/{episode}", movieViewerHandler)
	r.HandleFunc("/shows/{show}/{season}/{episode}/", movieViewerHandler)

	r.HandleFunc("/movies", movieHandler)
	r.HandleFunc("/movies/", movieHandler)

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	r.PathPrefix("/files/").Handler(http.StripPrefix("/files/", http.FileServer(http.Dir(mediaDirectory))))
	r.HandleFunc("/view/{movieOrShow}/{season}/{episode}", movieViewerHandler)
	r.HandleFunc("/view/{movieOrShow}", movieViewerHandler)
	r.HandleFunc("/", index)
	http.ListenAndServe(":8080", r)

}

func check(e error, w http.ResponseWriter) {

}
