package redis

import (
	// "fmt"
	"github.com/garyburd/redigo/redis"
)

func ConnectRedis(addr string) (state string, c redis.Conn, err error) {
	//Connect to Redis

	c, err = redis.Dial("tcp", addr)
	if err != nil {
		panic(err)
	}
	//Close if something went wrong
	//defer c.Close()
	state = "Connected"
	//username come from form login.html
	//fmt.Printf("Connected! ", c)
	return
}
