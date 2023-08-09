package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

const (
	// 메시지를 클라이언트에게 전송할 때의 최대 대기 시간
	writeWait = 10 * time.Second
	// 클라이언트가 pong 메시지를 보내야 하는 최대 시간
	pongWait = 60 * time.Second
	// ping 메시지 간격
	pingPeriod = (pongWait * 9) / 10
	// 메시지의 최대 크기
	maxMessageSize = 10000
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  4096,
	WriteBufferSize: 4096,
}

// Client는 연결된 클라이언트를 나타내는 구조체
type Client struct {
	// 클라이언트의 WebSocket 연결
	conn *websocket.Conn
	// 클라이언트가 속한 서버
	wsServer *WsServer
	// 메시지를 보내는 채널
	send chan []byte
}

// 새로운 클라이언트 객체를 생성하는 함수
func newClient(conn *websocket.Conn, wsServer *WsServer) *Client {
	return &Client{
		conn:     conn,
		wsServer: wsServer,
		send:     make(chan []byte, 256),
	}
}

// 클라이언트로부터 들어오는 메시지를 읽는 데 사용되는 함수
func (client *Client) readPump() {
	defer func() {
		client.disconnect()
	}()

	client.conn.SetReadLimit(maxMessageSize)
	client.conn.SetReadDeadline(time.Now().Add(pongWait))
	client.conn.SetPongHandler(func(string) error { client.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })

	// 클라이언트로부터 메시지를 계속 읽습니다.
	for {
		_, jsonMessage, err := client.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("unexpected close error: %v", err)
			}
			break
		}

		client.wsServer.broadcast <- jsonMessage
	}

}

// 서버에서 클라이언트로 메시지를 전송하는 함수
func (client *Client) writePump() {
	// 주어진 pingPeriod 간격으로 시간 틱을 생성하는 새로운 티커를 초기화
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		client.conn.Close()
	}()
	for {
		select {
		// 클라이언트가 전송할 메시지를 client.send 채널에서 기다립니다.
		// ok은 채널이 열려 있으면 true, 닫히면 false입니다.
		case message, ok := <-client.send:
			// 메시지를 클라이언트로 작성할 때의 최대 대기 시간을 설정
			client.conn.SetWriteDeadline(time.Now().Add(writeWait))
			// 채널이 닫힌 경우 (ok가 false인 경우) WebSocket 연결을 닫습니다.
			if !ok {
				client.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			// WebSocket 연결에 메시지를 작성하기 위한 writer를 생성
			w, err := client.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			// client.send 채널에 여러 메시지가 대기 중인 경우, 모든 메시지를 순서대로 보냅니다.
			n := len(client.send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-client.send)
			}

			if err := w.Close(); err != nil {
				return
			}

		// 티커가 설정된 시간 간격마다 발생하는 코드
		// 주기적으로 Ping 메시지를 보내어 클라이언트의 연결 상태를 확인
		case <-ticker.C:
			client.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := client.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

// 클라이언트를 연결 해제하는 함수
func (client *Client) disconnect() {
	client.wsServer.unregister <- client
	close(client.send)
	client.conn.Close()
}

// 새로운 클라이언트 WebSocket 연결을 처리하는 함수
func ServeWs(wsServer *WsServer, w http.ResponseWriter, r *http.Request) {

	// 일반 HTTP 연결을 WebSocket 연결로 업그레이드
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	// 새로운 Client 객체를 생성
	client := newClient(conn, wsServer)

	// 클라이언트의 메시지를 읽고 쓰는 데 사용되는 고루틴을 시작
	go client.writePump()
	go client.readPump()

	// 메인 서버 루프에 이 클라이언트를 등록하도록 요청
	wsServer.register <- client
}
