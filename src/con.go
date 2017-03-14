package main

import (
	"encoding/json"
	"github.com/garyburd/redigo/redis"
	"github.com/gorilla/websocket"
	"log"
)

var RedisServer bytes.Buffer

// WebSocket Message
//
// Used for JSON conversion.
// action = SUBSCRIBE|UNSUBSCRIBE|PUBLISH
// channel = Redis channel
// data = Message to be sent
type WSMessage struct {
	Action  string `json:"action"`
	Channel string `json:"channel"`
	Data    string `json:"data"`
}

// WebSocket Connection
//
// Handles incoming and outcoming websocket data by communicating
// with Redis via its PubSub commands.
type WSCon struct {
	socket    *websocket.Conn
	publish   *redis.PubSubConn
	subscribe *redis.PubSubConn
}

// Register at WSServer and connect to Redis.
func (wsc *WSCon) Initialize() {
	wss.register <- wsc
	wsc.publish = wsc.MakeRedisConnection()
	wsc.subscribe = wsc.MakeRedisConnection()
}

// Unregister from WSServer and disconnect from Redis.
func (wsc *WSCon) Uninitialize() {
	log.Println("Uninitialize")
	wss.unregister <- wsc

	wsc.publish.Close()
	wsc.subscribe.Close()
}

// Read from WebSocket
func (wsc *WSCon) ReadWebSocket() {
	for {
		var json_data []byte
		var message WSMessage

		// Receive data from WebSocket
		err := websocket.Message.Receive(wsc.socket, &json_data)
		if err != nil {
			return
		}

		// Parse JSON data
		err = json.Unmarshal(json_data, &message)
		if err != nil {
			return
		}
		switch message.Action {
		case "SUBSCRIBE":
			wsc.subscribe.Subscribe(message.Channel)
		case "UNSUBSCRIBE":
			wsc.subscribe.Unsubscribe(message.Channel)
		case "PUBLISH":
			wsc.publish.Conn.Do("PUBLISH", message.Channel, message.Data)
		}
	}
}

// Proxy incoming data from Redis to WebSocket.
func (wsc *WSCon) ProxyRedisSubscribe() {
	for {
		switch reply := wsc.subscribe.Receive().(type) {
		case redis.Message:
			message := WSMessage{"PUBLISH", reply.Channel, string(reply.Data)}
			json_data, err := json.Marshal(message)
			if err == nil {
				websocket.Message.Send(wsc.socket, string(json_data))
			}
		case redis.Subscription:
			message := WSMessage{"SUBSCRIBE", reply.Channel, ""}
			json_data, err := json.Marshal(message)
			if err == nil {
				websocket.Message.Send(wsc.socket, string(json_data))
			}
		case error:
			return
		}
	}
}

// Establish a connection to Redis via Redigo PubSubConn
func (wsc *WSCon) MakeRedisConnection() *redis.PubSubConn {
	RedisServer.WriteString("127.0.0.1")
	RedisServer.WriteString(":")
	RedisServer.WriteString("6379")
	stateRedis, c, errRedis := redis.ConnectRedis(RedisServer.String())
	defer c.Close()
	if errRedis != nil {
		panic(errRedis)
	}
	fmt.Println("status redis: ", stateRedis)
	return &redis.PubSubConn{c}
}

func handleWSConnection(socket *websocket.Conn) {
	wsc := &WSCon{socket: socket}
	defer wsc.Uninitialize()
	wsc.Initialize()
	go wsc.ProxyRedisSubscribe()
	wsc.ReadWebSocket()
}
