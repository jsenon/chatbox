package webserver

import (
	"fmt"
	"html/template"
	"net/http"
)

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
	username := req.FormValue("name")
	fmt.Println("Username:", username)
	Render(res, "src/templates/index.html", data)
}

func Error(res http.ResponseWriter, req *http.Request) {
}
