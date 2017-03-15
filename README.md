# ChatBox


Build CLI Chat Server based on redis

### Prerequisite

You need to have:

* Go 1.8
* Go Environment properly set
* Redis Docker image

### Installation

Start your redis

```sh
docker run -p 127.0.0.1:6379:6379 --name chat-redis -t redis
```
At this stage only CLI is implemented

Launch with your user ie. John
```sh
go run src/chatbox.go John
```

Launch with another user ie. Duff
```sh
go run src/chatbox.go Duff
```

### Access

CLI

### Todo

Construct Websocket



