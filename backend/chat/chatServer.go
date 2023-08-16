package main

type WsServer struct {
	clients    map[*Client]bool
	register   chan *Client
	unregister chan *Client
	broadcast  chan []byte
	rooms      map[*Room]bool
}

func NewWebsocketServer() *WsServer {
	return &WsServer{
		// 현재 연결된 모든 WebSocket 클라이언트를 저장하는 맵
		clients: make(map[*Client]bool),
		// 클라이언트 등록 채널
		register: make(chan *Client),
		// 클라이언트 등록 해제 채널
		unregister: make(chan *Client),
		// 브로드캐스트 채널
		broadcast: make(chan []byte),
		// 현재 생성된 모든 채팅방을 저장하는 맵
		rooms: make(map[*Room]bool),
	}
}

// 각각의 메시지에 따라 해당 클라이언트를 등록하거나 등록 해제하는 작업을 수행합니다.
// 또한, 브로드캐스트 메시지를 받으면 모든 클라이언트에게 메시지를 전송합니다.
func (server *WsServer) Run() {
	for {
		select {
		// 클라이언트 등록 채널에서 메시지를 기다립니다.
		case client := <-server.register:
			server.registerClient(client)
		// 클라이언트 등록 해제 채널에서 메시지를 기다립니다.
		case client := <-server.unregister:
			server.unregisterClient(client)
		// 브로드캐스트 채널에서 메시지를 기다립니다.
		case message := <-server.broadcast:
			server.broadcastToClients(message)
		}
	}
}

// 주어진 클라이언트를 clients 맵에 추가하는 함수
func (server *WsServer) registerClient(client *Client) {
	server.clients[client] = true
}

// 주어진 클라이언트를 clients 맵에서 제거하는 함수
func (server *WsServer) unregisterClient(client *Client) {
	if _, ok := server.clients[client]; ok {
		delete(server.clients, client)
	}
}

// 주어진 메시지를 모든 클라이언트에게 전송하는 함수
func (server *WsServer) broadcastToClients(message []byte) {
	for client := range server.clients {
		client.send <- message
	}
}

// 주어진 이름의 채팅방을 찾아서 반환하는 함수
func (server *WsServer) findRoomByName(name string) *Room {
	var foundRoom *Room
	for room := range server.rooms {
		if room.GetName() == name {
			foundRoom = room
			break
		}
	}

	return foundRoom
}

// 주어진 이름의 채팅방이 존재하면 해당 채팅방을 반환하고, 존재하지 않으면 새로운 채팅방을 생성하여 반환하는 함수
func (server *WsServer) createRoom(name string) *Room {
	room := NewRoom(name)
	go room.RunRoom()
	server.rooms[room] = true

	return room
}