package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"log"
	"net/http"
	"webserver"
)

// TO DO
// Build Docker File
// Construct and configure Redis
// Add function webserver
//

// TO FIX

func main() {

	//Connect to Redis
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		panic(err)
	}
	//Close if something went wrong
	defer c.Close()

	//username come from form login.html

	fmt.Printf("Connected! ", c)

	http.HandleFunc("/login", webserver.Login)
	http.HandleFunc("/mychat", webserver.Index)
	http.HandleFunc("/room", webserver.Room)

	//Handle URL ERROR
	http.HandleFunc("/", webserver.Error)
	// Init WebServer

	if err := http.ListenAndServe(":10000", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}

}
