package controllers

import (
	"net/http"
	"lenslocked.com/views"
)

type Users struct{
	// store all the templates we will be using to render users page 
	Templates struct {
		New views.Template
	}
}


// method to replace the signup page 
func (u Users) New(w http.ResponseWriter, r *http.Request) {
	u.Templates.New.Execute(w, nil)
}




