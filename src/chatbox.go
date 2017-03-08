package main

import (
	"fmt"
	"log"
	"net/http"
	"redis"
	"webserver"
)

// TO DO
// Build Docker File
// Construct and configure Redis
// Add function webserver
//

// TO FIX

var RedisServer = flag.String("127.0.0.1", ":6379", "")

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
	fmt.Println("connec chain", RediServer)
	c, err := redis.ConnectRedis(RedisServer)
	if err != nil {
		panic(err)
	}
	fmt.Println("status redis: ", c)

	// username.Name = req.FormValue("name")
	// userkey := "online." + username.Name
	// val, err := c.Do("SET", userkey, username.Name, "NX", "EX", "120")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if val == nil {
		fmt.Println("User already online")
		os.Exit(1)
	}

}
