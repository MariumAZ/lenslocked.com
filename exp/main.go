package main

import (
	"html/template"
	"os"
)
type User struct {
	Name string
	Bio string
	Age int
}

func main() {
	// t pointer to template
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		// no easyv way out to recover
		panic(err)
	}
	// create some data to put on our template
    user := User{
		Name : "Myriam",
		Bio: `<script>alert("Haha you have been h4x0r3d!");</script>`,
		Age: 123,
	}

	// execute template using this data
	err = t.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}

}
