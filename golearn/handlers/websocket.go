package handlers

import (
	"log"
	"net/http"

	ws "golearn/websocket"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func HandleWebSocket(hub *ws.Hub) gin.HandlerFunc {
	return func(c *gin.Context) {
		roomID := c.Param("roomId")
		userID := c.GetUint("user_id")
		username := c.Query("username")
		if username == "" {
			username = "Anonymous"
		}

		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			log.Printf("WebSocket upgrade error: %v", err)
			return
		}

		client := &ws.Client{
			Hub:      hub,
			Conn:     conn,
			Send:     make(chan ws.Message, 256),
			RoomID:   roomID,
			UserID:   userID,
			Username: username,
		}

		hub.Register <- client

		go writePump(client, conn)
		go readPump(client, conn, hub)
	}
}

func readPump(client *ws.Client, conn *websocket.Conn, hub *ws.Hub) {
	defer func() {
		hub.Unregister <- client
		conn.Close()
	}()

	for {
		var msg ws.Message
		if err := conn.ReadJSON(&msg); err != nil {
			break
		}
		msg.RoomID = client.RoomID
		msg.UserID = client.UserID
		msg.Username = client.Username
		msg.Type = "message"

		hub.Broadcast <- msg
	}
}

func writePump(client *ws.Client, conn *websocket.Conn) {
	defer conn.Close()

	for msg := range client.Send {
		if err := conn.WriteJSON(msg); err != nil {
			break
		}
	}
}
