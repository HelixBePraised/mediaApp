package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strings"

	"os"

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

func index(w http.ResponseWriter, req *http.Request) {
	_ = req.URL.Path
	p := Page{
		Title: "Home",
	}

	err := tmpl.ExecuteTemplate(w, "home.gohtml", p)
	check(err, w)

}

func movieHandler(w http.ResponseWriter, req *http.Request) {
	p := Page{
		Title:             "Movies",
		MediaTitleAndLink: movies,
	}

	err := tmpl.ExecuteTemplate(w, "movieIndex.gohtml", p)
	check(err, w)

}

func showsHandler(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	var p Page
	var title string
	if vars["show"] == "" {
		var m = map[string]string{}

		for show := range shows {
			m[show] = "/shows/" + show
		}

		p = Page{
			Title:             "Shows",
			MediaTitleAndLink: m,
		}

	} else if vars["season"] == "" {
		var m = map[string]string{}

		for season := range shows[vars["show"]] {
			if season == vars["show"] {
				continue
			}
			m[season] = "/shows/" + vars["show"] + "/" + season
		}

		title = vars["show"]

		p = Page{
			Title:             title,
			MediaTitleAndLink: m,
		}

	} else if vars["episode"] == "" {
		var m = map[string]string{}

		for episode := range shows[vars["show"]][vars["season"]] {
			m[episode] = "/view/" + vars["show"] + "/" + vars["season"] + "/" + episode
		}

		title = vars["show"] + " | " + vars["season"]

		p = Page{
			Title:             title,
			MediaTitleAndLink: m,
		}
	}

	err := tmpl.ExecuteTemplate(w, "movieIndex.gohtml", p)
	check(err, w)
}

func movieViewerHandler(w http.ResponseWriter, req *http.Request) {
	var url, title string
	vars := mux.Vars(req)

	if vars["season"] == "" {
		url = "/files/movies/" + vars["movieOrShow"]
		title = vars["movieOrShow"]
	} else {
		url = "/files/shows/" + vars["movieOrShow"] + "/" + vars["season"] + "/" + vars["episode"]
		title = vars["movieOrShow"]
	}

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
