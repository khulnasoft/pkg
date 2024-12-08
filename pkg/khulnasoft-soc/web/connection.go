package web

import (
	"time"

	"github.com/gorilla/websocket"
)

type Connection struct {
	websocket    *websocket.Conn
	lastPingTime time.Time
	ip           string
}

func NewConnection(wsConn *websocket.Conn, ip string) *Connection {
	conn := &Connection{
		websocket: wsConn,
		ip:        ip,
	}
	conn.UpdatePingTime()
	return conn
}

func (connection *Connection) IsAuthorized(kind string) bool {
	return true
}

func (connection *Connection) UpdatePingTime() {
	connection.lastPingTime = time.Now()
}
