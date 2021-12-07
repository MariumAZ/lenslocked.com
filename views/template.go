package views

import (
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"net/http"
)

// goal
// wraps around html we have been using

type Template struct {
	HTMLTpl *template.Template

}

func Must(t Template, err error) Template {
	if err != nil {
		panic(err)
	}
	return t
}

func ParseFS(fs fs.FS, pattern string)(Template, error){
    tpl, err := template.ParseFS(fs, pattern)
	if err != nil {
		return Template{}, fmt.Errorf("parse template: %w", err)
	}
	return Template{
		HTMLTpl: tpl,
	}, nil
}

func Parse(path string) (Template, error) {
	tpl, err := template.ParseFiles(path)
	if err != nil {
		return Template{}, fmt.Errorf("parse template: %w", err)
	}
	return Template{
		HTMLTpl: tpl,
	}, nil
}


//take sin the data that is going to be rendered 
// and because that's what template.execute takes in
// we will assume that our template is already parsed 
func (t Template) Execute(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-type", "text/html")
	err := t.HTMLTpl.Execute(w, nil) // writing the output to w 
	if err != nil {
		log.Printf("executing template: %v", err)
		http.Error(w, "there ws an error executing the template", http.StatusInternalServerError)
		return
	}
}