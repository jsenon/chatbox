package main

import (
	// "fmt"
	"net/http"
	// "os"
	// "bufio"
	"bytes"
	// redigo "github.com/garyburd/redigo/redis"
	"log"
	// "os"
	// "redis"
	// "time"
	"webserver"
)

// TO DO
// Add function webserver
//

// TO FIX

var RedisServer bytes.Buffer

// For testPurpose before retrieve it from webserver
var username string

func main() {

	// Web Part
	r := newRoom()

	http.HandleFunc("/login", webserver.Login)
	http.HandleFunc("/mychat", webserver.Index)
	http.HandleFunc("/room", r)
	go r.run()
	//Handle URL ERROR
	http.HandleFunc("/", webserver.Error)
	// Init WebServer
	if err := http.ListenAndServe(":10000", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}

}
