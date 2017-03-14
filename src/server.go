package main

type WSServer struct {
	con        map[*WSCon]bool
	register   chan *WSCon
	unregister chan *WSCon
}

func (wss *WSServer) Run() {
	for {
		select {
		case wsc := <-wss.register:
			wss.connection[wsc] = true
		case wsc := <-wss.unregister:
			delete(wss.con, wsc)
		}
	}
}

var wss = WSServer{
	con:        make(map[*WSCon]bool),
	register:   make(chan *WSCon),
	unregister: make(chan *WSCon),
}

// Wrapper function for wss.Run
func runWSServer() {
	wss.Run()
}
