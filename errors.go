package main

import (
	"fmt"
	"net/http"
)

//Fine for now, possibly want to do more with errors
func check(err error, w http.ResponseWriter) {
	if err != nil {
		fmt.Println(err)
		fmt.Fprintf(w, "%s", err)
	}
}