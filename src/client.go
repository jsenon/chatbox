package main

import (
	"github.com/gorilla/websocket"
)

// client represents a single chatting user.
type client struct {

	// socket is the web socket for this client.
	socket *websocket.Conn

	// send is a channel on which messages are sent.
	send chan []byte

	// room is the room this client is chatting in.
	room *room
}

func (c *client) read() {
	defer c.socket.Close()
	for {
		_, msg, err := c.socket.ReadMessage()
		if err != nil {
			return
		}
		c.room.forward <- msg
	}
}

func (c *client) write() {
	defer c.socket.Close()
	for msg := range c.send {
		err := c.socket.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			return
		}
	}
}

// DRAFT TO BE DEFINED

// RedisServer.WriteString("127.0.0.1")
// RedisServer.WriteString(":")
// RedisServer.WriteString("6379")

// fmt.Println("connec chain", RedisServer.String())

// // Call Redis connection
// stateRedis, c, errRedis := redis.ConnectRedis(RedisServer.String())
// defer c.Close()

// if errRedis != nil {
// 	panic(errRedis)
// }
// fmt.Println("status redis: ", stateRedis)

// username := John
// userkey := "online." + username
// // EX seconds -- Set the specified expire time, in seconds.
// // PX milliseconds -- Set the specified expire time, in milliseconds.
// // NX -- Only set the key if it does not already exist.
// // XX -- Only set the key if it already exist.
// // We set a key online.John with user John only if key is not present and set expiration to 30s
// val, err := c.Do("SET", userkey, username, "NX", "EX", "30")
// // Handle Error
// if err != nil {
// 	fmt.Println(err)
// 	panic(err)
// }
// // Handle Key already present, username already online
// if val == nil {
// 	fmt.Println("User already online")
// 	// To be redirected on future login page
// 	panic(val)
// }
// // fmt.Println("Log redis", c)
// // We add the specified members to the set stored at key users
// val, err = c.Do("SADD", "users", username)
// if err != nil {
// 	fmt.Println(err)
// 	panic(err)
// }
// // Handle user already member
// if val == nil {
// 	fmt.Println("User already stored")
// 	panic(val)
// }

// chatExit := false

// for !chatExit {
// 	select {
// 	case msg := <-subChan:
// 		fmt.Println(msg)
// 	// case <-tickerChan:
// 	// 	val, err = c.Do("SET", userkey, username, "XX", "EX", "120")
// 	// 	if err != nil || val == nil {
// 	// 		fmt.Println("Heartbeat set failed")
// 	// 		chatExit = true
// 	// 	}
// 	case line := <-sayChan:
// 		// Comment exit cammand line for server version, we don't want to kill program
// 		if line == "/exit" {
// 			chatExit = true
// 		} else if line == "/who" {
// 			names, _ := redigo.Strings(c.Do("SMEMBERS", "users"))
// 			for _, name := range names {
// 				fmt.Println(name)
// 			}
// 		} else {
// 			c.Do("PUBLISH", "messages", username+":"+line)
// 		}
// 	default:
// 		time.Sleep(100 * time.Millisecond)
// 	}
// }

// // Logout User
// // Clean all keys
// // Publish left message
// c.Do("DEL", userkey)
// c.Do("SREM", "users", username)
// c.Do("PUBLISH", "messages", username+" has left")
// // TEST PURPOSE ONLY
// // A ticker will let us update our presence on the Redis server
// // tickerChan := time.NewTicker(time.Second * 60).C
