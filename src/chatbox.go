package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	// "os"
	// "redis"
	"webserver"
)

// TO DO
// Build Docker File
// Construct and configure Redis
// Add function webserver
//

// TO FIX

var RedisServer = flag.String("ipaddr", "127.0.0.1", ":6379")

func main() {

	fmt.Println("connec chain", *RedisServer)

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
	fmt.Println("connec chain", RedisServer)
	//c, err := redis.ConnectRedis(RedisServer)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("status redis: ", c)

	// username.Name = req.FormValue("name")
	// userkey := "online." + username.Name
	// val, err := c.Do("SET", userkey, username.Name, "NX", "EX", "120")

}
