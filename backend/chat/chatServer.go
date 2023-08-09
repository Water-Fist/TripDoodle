package main

type WsServer struct {
	clients    map[*Client]bool
	register   chan *Client
	unregister chan *Client
}

func NewWebsocketServer() *WsServer {
	return &WsServer{
		// 현재 연결된 모든 WebSocket 클라이언트를 저장하는 맵
		clients: make(map[*Client]bool),
		// 클라이언트 등록 채널
		register: make(chan *Client),
		// 클라이언트 등록 해제 채널
		unregister: make(chan *Client),
	}
}

// 각각의 메시지에 따라 해당 클라이언트를 등록하거나 등록 해제하는 작업을 수행합니다.
func (server *WsServer) Run() {
	for {
		select {
		// 클라이언트 등록 채널에서 메시지를 기다립니다.
		case client := <-server.register:
			server.registerClient(client)
		// 클라이언트 등록 해제 채널에서 메시지를 기다립니다.
		case client := <-server.unregister:
			server.unregisterClient(client)
		}
	}
}

// 주어진 클라이언트를 clients 맵에 추가하는 메서드
func (server *WsServer) registerClient(client *Client) {
	server.clients[client] = true
}

// 주어진 클라이언트를 clients 맵에서 제거하는 메서드
func (server *WsServer) unregisterClient(client *Client) {
	if _, ok := server.clients[client]; ok {
		delete(server.clients, client)
	}
}
