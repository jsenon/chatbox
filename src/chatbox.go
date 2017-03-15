package main

import (
	"bufio"
	"bytes"
	"fmt"
	redigo "github.com/garyburd/redigo/redis"
	"os"
	"redis"
	"time"
)

// TO DO
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
		fmt.Println("Usage: chatbox username")
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
	// fmt.Println("Log redis", c)
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

	// Subscribing message
	subChan := make(chan string)
	go func() {
		stateRedisSub, subconn, errRedisSub := redis.ConnectRedis(RedisServer.String())
		if errRedisSub != nil {
			fmt.Println(errRedisSub)
			fmt.Println(stateRedisSub)
			os.Exit(1)
		}
		defer subconn.Close()

		psc := redigo.PubSubConn{Conn: subconn}
		// Subscribe to channel messages
		psc.Subscribe("messages")
		for {
			switch v := psc.Receive().(type) {
			case redigo.Message:
				subChan <- string(v.Data)
			case redigo.Subscription:
				// We don't need to listen to subscription messages,
			case error:
				return
			}
		}
	}()

	// Sending message
	sayChan := make(chan string)
	go func() {
		prompt := username + ">"
		bio := bufio.NewReader(os.Stdin)
		for {
			fmt.Print(prompt)
			line, _, err := bio.ReadLine()
			if err != nil {
				fmt.Println(err)
				sayChan <- "/exit"
				return
			}
			sayChan <- string(line)
		}

	}()

	c.Do("PUBLISH", "messages", username+" has joined")
	chatExit := false

	for !chatExit {
		select {
		case msg := <-subChan:
			fmt.Println(msg)
		// case <-tickerChan:
		// 	val, err = c.Do("SET", userkey, username, "XX", "EX", "120")
		// 	if err != nil || val == nil {
		// 		fmt.Println("Heartbeat set failed")
		// 		chatExit = true
		// 	}
		case line := <-sayChan:
			if line == "/exit" {
				chatExit = true
			} else if line == "/who" {
				names, _ := redigo.Strings(c.Do("SMEMBERS", "users"))
				for _, name := range names {
					fmt.Println(name)
				}
			} else {
				c.Do("PUBLISH", "messages", username+":"+line)
			}

		default:
			time.Sleep(100 * time.Millisecond)
		}
	}

	// Logout User
	// Clean all keys
	// Publish left message
	c.Do("DEL", userkey)
	c.Do("SREM", "users", username)
	c.Do("PUBLISH", "messages", username+" has left")
}
