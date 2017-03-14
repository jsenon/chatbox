package main

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"webserver"
)

// TO DO
// Add function webserver
//

// TO FIX

// For testPurpose before retrieve it from webserver
func main() {

	http.HandleFunc("/mychat", webserver.Index)
	// Handle URL ERROR
	http.HandleFunc("/", webserver.Error)
	// Init WebServer
	http.Handle("/ws", websocket.Handler(handleWSConnection))
	if err := http.ListenAndServe(":10000", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
