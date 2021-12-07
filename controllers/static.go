package controllers

import (
	"html/template"
	"net/http"

	"lenslocked.com/views"
)

func StaticHandler(tpl views.Template) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, nil)
	}
}	

func FAQ(tpl views.Template) http.HandlerFunc {
    questions := []struct{
		Question string
		Answer template.HTML
	}{
	{
		Question : "Is there a free version ?",
        Answer : "Yes sure !",
	},
	{
		Question : "Do you do other things than coding in life ?",
		Answer : "Bien Sur I do workout ;) ",

	},
    {
		Question: "How Can I get support",
		Answer: `Please feel free to Email us : <a href="mailto:myriam.azzouz@holbertonschool.com">myriam.azzouz@holbertonschool.com</a>`,
	},
}
    return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, questions)
	}
}

