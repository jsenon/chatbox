package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
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

	response, err := c.Do("AUTH", "YOUR_PASSWORD")

	if err != nil {
		panic(err)
	}

	fmt.Printf("Connected! ", response)
	defer c.Close()

}
