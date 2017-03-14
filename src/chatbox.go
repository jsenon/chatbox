package main

import (
	"bytes"
	"code.google.com/p/go.net/websocket"
	"log"
	"net/http"
)

// TO DO
// Add function webserver
//

// TO FIX

// For testPurpose before retrieve it from webserver
func main() {

	http.HandleFunc("/mychat", webserver.Index)
	go r.run()
	// Handle URL ERROR
	http.HandleFunc("/", webserver.Error)
	// Init WebServer
	http.Handle("/ws", websocket.Handler(handleWSConnection))
	if err := http.ListenAndServe(":10000", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
