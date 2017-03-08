package main

import (
	"fmt"
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

	http.HandleFunc("/login", webserver.Login)
	http.HandleFunc("/mychat", webserver.Index)
	http.HandleFunc("/room", webserver.Room)

	//Handle URL ERROR
	http.HandleFunc("/", webserver.Error)
	// Init WebServer

	if err := http.ListenAndServe(":10000", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}

	// Call Redis connection

}
