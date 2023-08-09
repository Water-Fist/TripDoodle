package main

import "fmt"

type Room struct {
	name       string
	clients    map[*Client]bool
	register   chan *Client
	unregister chan *Client
	broadcast  chan *Message
}

// 새로운 채팅방 생성
func NewRoom(name string) *Room {
	return &Room{
		name:       name,
		clients:    make(map[*Client]bool),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		broadcast:  make(chan *Message),
	}
}

func (room *Room) GetName() string {
	return room.name
}

func (room *Room) RunRoom() {
	for {
		select {
		// 클라이언트가 채팅방에 등록될 때 이 채널을 통해 알림을 받습니다.
		case client := <-room.register:
			room.registerClientInRoom(client)
		// 클라이언트가 채팅방에서 등록 해제될 때 이 채널을 통해 알림을 받습니다.
		case client := <-room.unregister:
			room.unregisterClientInRoom(client)
		// 클라이언트가 메시지를 보낼 때 이 채널을 통해 알림을 받습니다.
		case message := <-room.broadcast:
			room.broadcastToClientsInRoom(message.encode())
		}
	}
}

// 클라이언트를 채팅방에 등록하는 함수
func (room *Room) registerClientInRoom(client *Client) {
	room.notifyClientJoined(client)
	room.clients[client] = true
}

// 클라이언트를 채팅방에서 등록 해제하는 함수
func (room *Room) unregisterClientInRoom(client *Client) {
	if _, ok := room.clients[client]; ok {
		delete(room.clients, client)
	}
}

// 채팅방에 메시지를 브로드캐스팅하는 함수
func (room *Room) broadcastToClientsInRoom(message []byte) {
	for client := range room.clients {
		client.send <- message
	}
}

const welcomeMessage = "%s joined the room"

// 클라이언트가 채팅방에 들어오면 클라이언트에게 환영 메시지를 전송합니다.
func (room *Room) notifyClientJoined(client *Client) {
	message := &Message{
		Action:  SendMessageAction,
		Target:  room.name,
		Message: fmt.Sprintf(welcomeMessage, client.GetName()),
	}

	room.broadcastToClientsInRoom(message.encode())
}
