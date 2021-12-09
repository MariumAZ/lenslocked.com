package main

import (
	"fmt"
	_ "log"
	"net/http"
	"github.com/go-chi/chi/v5"
	"lenslocked.com/controllers"
	"lenslocked.com/templates"
	"lenslocked.com/views"
	_"path/filepath"
)

func notfound(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusNotFound)
	w.Header().Set("Content-type", "text/html")
	fmt.Fprint(w, "OOps 404 ! ")
}
func main() {
	// initialize a router
	r := chi.NewRouter()
	//views.Must(views.ParseFS(templates.FS,"home.gohtml"))

	tpl := views.Must(views.ParseFS(templates.FS,"home.gohtml", "tailwind.gohtml"))
	r.Get("/", controllers.StaticHandler(tpl))
   
    tpl = views.Must(views.ParseFS(templates.FS,"Me.gohtml", "tailwind.gohtml"))
	r.Get("/aboutMe", controllers.StaticHandler(tpl))

	tpl = views.Must(views.ParseFS(templates.FS,"contact.gohtml", "tailwind.gohtml"))
    r.Get("/contact", controllers.StaticHandler(tpl))

	tpl = views.Must(views.ParseFS(templates.FS,"faq.gohtml", "tailwind.gohtml"))
    r.Get("/faq", controllers.FAQ(tpl))
	
	// the handler function is an actual type
	r.NotFound(notfound)
	http.ListenAndServe(":3000", r)   
}


