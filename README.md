# ChatBox


Build Chat Server based on redis

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
At this stage only offline/online communication with redis is available

Simulate John
```sh
go run src/chatbox.go John
```

Simulate Duff
```sh
go run src/chatbox.go Duff
```

### Access

TBD

### Todo

Construct Docker Redis




