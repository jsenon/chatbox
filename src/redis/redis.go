package redis

import (
	"github.com/garyburd/redigo/redis"
)

// Need addr IP:Port in entry
// Return state string has connected if OK
// Return c has connector to our redis server
// Return err if error during connection

func ConnectRedis(addr string) (state string, c redis.Conn, err error) {
	//Connect to Redis
	c, err = redis.Dial("tcp", addr)
	if err != nil {
		panic(err)
	}
	state = "Connected"
	return
}
