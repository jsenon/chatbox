package main

import (
	redigo "github.com/garyburd/redigo/redis"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"trace"
)

type room struct {
	forward chan []byte
	join    chan *client
	leave   chan *client
	client  map[*client]bool
	tracer  trace.Tracer
}

func newRoom() *room {
	return &room{
		forward: make(chan []byte),
		join:    make(chan *client),
		leave:   make(chan *client),
		clients: make(map[*client]bool),
		tracer:  trace.Off(),
	}
}

func (r *room) run() {

	RedisServer.WriteString("127.0.0.1")
	RedisServer.WriteString(":")
	RedisServer.WriteString("6379")

	fmt.Println("connec chain", RedisServer.String())

	// Call Redis connection
	stateRedis, c, errRedis := redis.ConnectRedis(RedisServer.String())
	defer c.Close()

	select {
	case client := <-r.join:
		// join
		r.clients[client] = true
		c.Do("PUBLISH", "messages", "Client has joined")
	case client := <-r.leave:
		//ciao
		delete(r.clients, client)
		close(client.send)
		c.Do("PUBLISH", "messages", "Client has left")
	case msg := <-r.forward:
		c.Do("PUBLISH", "messages", ":"+string(msg.Message))
	}
}

const (
	socketBufferSize  = 1024
	messageBufferSize = 256
)

var upgrader = &websocket.Upgrader{ReadBufferSize: socketBufferSize, WriteBufferSize: socketBufferSize}

func (r *room) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	socket, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Fatal("ServeHTTP:", err)
		return
	}

	authCookie, err := req.Cookie("auth")
	if err != nil {
		log.Fatal("Failed to get auth cookie:", err)
		return
	}
	client := &client{
		socket:   socket,
		send:     make(chan *message, messageBufferSize),
		room:     r,
		userData: objx.MustFromBase64(authCookie.Value),
	}
	r.join <- client
	defer func() { r.leave <- client }()
	go client.write()
	client.read()
}
