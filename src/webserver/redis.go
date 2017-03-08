package redis

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func ConnectRedis(ipname string, port float64) (state string) {
	//Connect to Redis
	
	c, err := redis.Dial("tcp", ipname:port)
	if err != nil {
		panic(err)
	}
	//Close if something went wrong
	defer c.Close()
	state = "Connected"
	return state
	//username come from form login.html
	fmt.Printf("Connected! ", c)
}
