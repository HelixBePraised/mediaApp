package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

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
