package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	// content type is like the key here 
	// smthing related to web applications : 
	// web requests it is not smthing related to go 
	// to return html pages 
	w.Header().Set("Content-Type", "text/html")
	// we are in the home page
	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1> Hello my Dear Guests </h1>")
	} else if r.URL.Path == "/contact" {
		fmt.Fprint(w, "To get in touch please send an email to <a href=\"mailto:myriam.azzouz@holbertonschool.com\"> support </a>.")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Ooops the page not available ! ")
	}
	}
func main() {
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":3000", nil)

}


