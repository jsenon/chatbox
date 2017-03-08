package webserver

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"html/template"
	"net/http"
	"os"
)

// Function for Rendering templates
// filename is relative path form where you run the bin

type User struct {
	Name string
	Age  int
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
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		panic(err)
	}
	//Close if something went wrong
	defer c.Close()

	username.Name = req.FormValue("name")

	userkey := "online." + username.Name

	val, err := c.Do("SET", userkey, username.Name, "NX", "EX", "120")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if val == nil {
		fmt.Println("User already online")
		os.Exit(1)
	}

	Render(res, "src/templates/index.html", data)
}

func Login(res http.ResponseWriter, req *http.Request) {
	Render(res, "src/templates/login.html", nil)
}

//Return User

func Room(res http.ResponseWriter, req *http.Request) {

}

func Error(res http.ResponseWriter, req *http.Request) {

}
