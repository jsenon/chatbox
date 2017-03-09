package main

import (
	"fmt"
	"net/http"
	// "os"
	"bytes"
	redigo "github.com/garyburd/redigo/redis"
	"os"
	"redis"
	"time"
	"webserver"
)

// TO DO
// Construct pu
// Add function webserver
//

// TO FIX

var RedisServer bytes.Buffer

// For testPurpose before retrieve it from webserver
var username string

func main() {

	RedisServer.WriteString("127.0.0.1")
	RedisServer.WriteString(":")
	RedisServer.WriteString("6379")
	// username = "John"

	// TEST PURPOSE
	// Username is retrieve through argument in command line
	if len(os.Args) != 2 {
		fmt.Println("Usage: goredchat username")
		os.Exit(1)
	}
	username := os.Args[1]

	fmt.Println("connec chain", RedisServer.String())

	// Call Redis connection
	stateRedis, c, errRedis := redis.ConnectRedis(RedisServer.String())
	defer c.Close()

	if errRedis != nil {
		panic(errRedis)
	}
	fmt.Println("status redis: ", stateRedis)
	// Create a key
	userkey := "online." + username
	// EX seconds -- Set the specified expire time, in seconds.
	// PX milliseconds -- Set the specified expire time, in milliseconds.
	// NX -- Only set the key if it does not already exist.
	// XX -- Only set the key if it already exist.
	// We set a key online.John with user John only if key is not present and set expiration to 30s
	val, err := c.Do("SET", userkey, username, "NX", "EX", "30")
	// Handle Error
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	// Handle Key already present, username already online
	if val == nil {
		fmt.Println("User already online")
		// To be redirected on future login page
		panic(val)
	}
	fmt.Println("Log redis", c)
	// We add the specified members to the set stored at key users
	val, err = c.Do("SADD", "users", username)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	// Handle user already member
	if val == nil {
		fmt.Println("User already stored")
		panic(val)
	}
	// Print who is connected
	// Set timer 2s
	duration := time.Duration(2) * time.Second
	// TEST PURPOSE
	// Check online/offline function loop 60s
	for i := 0; i < 30; i++ {
		// Force John lgout
		if username == "John" && i == 15 {
			c.Do("DEL", userkey)
			c.Do("SREM", "users", username)
			c.Do("PUBLISH", "messages", username+" has left")
			os.Exit(1)
		}
		// Who is online
		fmt.Println("-----")
		names, _ := redigo.Strings(c.Do("SMEMBERS", "users"))
		for _, name := range names {
			fmt.Println("Users online: ", name)
		}
		time.Sleep(duration)
	}
	// Logout User
	// Clean all keys
	// Publish left message
	c.Do("DEL", userkey)
	c.Do("SREM", "users", username)
	c.Do("PUBLISH", "messages", username+" has left")
	// TEST PURPOSE ONLY
	// A ticker will let us update our presence on the Redis server
	// tickerChan := time.NewTicker(time.Second * 60).C

	// Web Part
	//Handle Func
	http.HandleFunc("/login", webserver.Login)
	http.HandleFunc("/mychat", webserver.Index)
	http.HandleFunc("/room", webserver.Room)
	//Handle URL ERROR
	http.HandleFunc("/", webserver.Error)
	// Init WebServer
	// if err := http.ListenAndServe(":10000", nil); err != nil {
	// 	log.Fatal("ListenAndServe:", err)
	// }
}
