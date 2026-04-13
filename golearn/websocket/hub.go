package websocket

import (
	"sync"
)

type Message struct {
	RoomID   string `json:"room_id"`
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	Content  string `json:"content"`
	Type     string `json:"type"`
}

type Client struct {
	Hub      *Hub
	Conn     interface{ WriteJSON(v interface{}) error; ReadJSON(v interface{}) error; Close() error }
	Send     chan Message
	RoomID   string
	UserID   uint
	Username string
}

type Hub struct {
	Rooms      map[string]map[*Client]bool
	Broadcast  chan Message
	Register   chan *Client
	Unregister chan *Client
	mu         sync.RWMutex
}

func NewHub() *Hub {
	return &Hub{
		Rooms:      make(map[string]map[*Client]bool),
		Broadcast:  make(chan Message),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.mu.Lock()
			if h.Rooms[client.RoomID] == nil {
				h.Rooms[client.RoomID] = make(map[*Client]bool)
			}
			h.Rooms[client.RoomID][client] = true
			h.mu.Unlock()

			// Notify room
			h.Broadcast <- Message{
				RoomID:   client.RoomID,
				UserID:   client.UserID,
				Username: client.Username,
				Content:  client.Username + " joined the room",
				Type:     "system",
			}

		case client := <-h.Unregister:
			h.mu.Lock()
			if clients, ok := h.Rooms[client.RoomID]; ok {
				if _, exists := clients[client]; exists {
					delete(clients, client)
					close(client.Send)
					if len(clients) == 0 {
						delete(h.Rooms, client.RoomID)
					}
				}
			}
			h.mu.Unlock()

			// Notify room
			h.Broadcast <- Message{
				RoomID:   client.RoomID,
				UserID:   client.UserID,
				Username: client.Username,
				Content:  client.Username + " left the room",
				Type:     "system",
			}

		case msg := <-h.Broadcast:
			h.mu.RLock()
			if clients, ok := h.Rooms[msg.RoomID]; ok {
				for client := range clients {
					select {
					case client.Send <- msg:
					default:
						close(client.Send)
						delete(clients, client)
					}
				}
			}
			h.mu.RUnlock()
		}
	}
}
