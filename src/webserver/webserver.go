package webserver

import (
	"fmt"
	"html/template"
	"net/http"
)

// Function for Rendering templates
// filename is relative path form where you run the bin

type User struct {
	Name string
}

var username User

func Render(w http.ResponseWriter, filename string, data interface{}) {
	tmpl, err := template.ParseFiles(filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func Index(res http.ResponseWriter, req *http.Request) {
	data := struct {
		Title    string
		Myapikey string
	}{
		Title:    "",
		Myapikey: "",
	}

	req.ParseForm()
	username.Name = req.FormValue("name")
	fmt.Println("Username:", username)
	Render(res, "src/templates/index.html", data)
}

func Login(res http.ResponseWriter, req *http.Request) {
	Render(res, "src/templates/login.html", nil)
}

func Error(res http.ResponseWriter, req *http.Request) {

}
