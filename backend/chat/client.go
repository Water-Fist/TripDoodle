package chat

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  4096,
	WriteBufferSize: 4096,
}

// 서버에 웹소켓 클라이언트를 나타냄
type Client struct {
	//웹소켓 연결
	conn *websocket.Conn
}

// 새 클라이언트를 만듬
func newClient(conn *websocket.Conn) *Client {
	return &Client{
		conn: conn,
	}
}

// 클라이언트 요청으로부터 웹소켓 요청을 핸들링
func ServeWs(w http.ResponseWriter, r *http.Request) {

	//WebSocket 프로토콜에 대한 HTTP 서버 연결을 업그레이드하는 데 사용
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	client := newClient(conn)

	fmt.Println("New Client joined the hub!")
	fmt.Println(client)
}
