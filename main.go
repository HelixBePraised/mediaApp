package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type HomePage struct {
	Title string
}

func main() {
	http.HandleFunc("/", index)
	//http.HandleFunc("/movies/", movieIndex)
	http.ListenAndServe(":8080", nil)

}

func index(w http.ResponseWriter, req *http.Request) {
	//w.Header().Set("Content-Type", "text/html")
	p := HomePage{Title: "Home"}

	t, err := template.ParseFiles("./temp/home.html")

	if err != nil {
		fmt.Println(err)
	}

	err = t.Execute(w, p)

	if err != nil {
		fmt.Println(err)
		fmt.Fprintf(w, "%s", err)
	}
}

// func movieIndex(w http.ResponseWriter, req *http.Request)  {
// 	w.Header().Set("Content-type", "text/html")
//
// }
