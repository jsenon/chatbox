package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"net/http"
)

// TO DO
// Build Docker File
// Construct and configure Redis
//

// TO FIX

func main() {

	//Connect to Redis
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		panic(err)
	}
	if err != nil {
		panic(err)
	}

	response, err := c.Do("AUTH", "YOUR_PASSWORD")

	fmt.Printf("Connected! ", response)
	defer c.Close()

	http.HandleFunc("/mychat", webserver.Index)
	http.HandleFunc("/room", webserver.Room)

	//Handle URL ERROR
	http.HandleFunc("/", webserver.Error)
	// Init WebServer

	if err := http.ListenAndServe(":10000", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}

}
