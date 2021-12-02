package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)


func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html")
	fmt.Fprint(w, "<h1>Welcome HOME!</h1>")
}	

func contact(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-type", "text/html")
	fmt.Fprint(w, "To get in touch please send an email to <a href=\"mailto:myriam.azzouz@holbertonschool.com\"> support </a>.")
}

func faq(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-type", "text/html")
	fmt.Fprint(w, "<h1>FAQ</h1>")
	fmt.Fprint(w, "1/ Who should Icontact in case I got 404 ?")
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


