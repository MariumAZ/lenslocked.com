package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"github.com/gorilla/mux"
	"lenslocked.com/views"
)

// we could have passed an io.writer but http.error later on
// excpects an http.response writer
func executeTemplate(w http.ResponseWriter, filepath string) {
	w.Header().Set("Content-type", "text/html")
	tpl, err := views.Parse(filepath)
	//err = tmp.Execute(w, nil) // writing the output to w 
	if err != nil {
		log.Printf("executing template: %v", err)
		http.Error(w, "there ws an error executing the template", http.StatusInternalServerError)
		return
	}
	tpl.Execute(w, nil)
	//viewTpl := views.Template{
	//	HTMLTpl: tpl,
	//}
	//viewTpl.Execute(w, nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html")
	tplpath := filepath.Join("templates", "home.gohtml")
	executeTemplate(w, tplpath)
}	

func contact(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-type", "text/html")
	tplpath := filepath.Join("templates", "contact.gohtml")
	executeTemplate(w, tplpath)
}

func faq(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-type", "text/html")
	tplpath := filepath.Join("templates", "faq.gohtml")
    executeTemplate(w, tplpath)
}
func notfound(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusNotFound)
	w.Header().Set("Content-type", "text/html")
	fmt.Fprint(w, "OOps 404 ! ")
}
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	r.HandleFunc("/FAQ", faq)
	// the handler function is an actual type
	r.NotFoundHandler = http.HandlerFunc(notfound)
	http.ListenAndServe(":3000", r)
	    
}


