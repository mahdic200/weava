package Websocket

import (
	"sync"

	"github.com/gofiber/contrib/websocket"
)

type Client struct {
	ID   string
	Conn *websocket.Conn
}

type Room struct {
	Clients   map[*Client]bool
	ClientsMu sync.Mutex
}

type Manager struct {
	Rooms   map[string]*Room
	RoomsMu sync.RWMutex
}

func NewManager() *Manager {
	return &Manager{
		Rooms: make(map[string]*Room),
	}
}

func (m *Manager) CreateRoom(roomID string) *Room {
	m.RoomsMu.Lock()
	defer m.RoomsMu.Unlock()

	room := &Room{
		Clients: make(map[*Client]bool),
	}
	m.Rooms[roomID] = room
	return room
}

func (m *Manager) GetRoom(roomID string) *Room {
	m.RoomsMu.RLock()
	defer m.RoomsMu.RUnlock()
	return m.Rooms[roomID]
}

func (r *Room) AddClient(client *Client) {
	r.ClientsMu.Lock()
	defer r.ClientsMu.Unlock()
	r.Clients[client] = true
}

func (r *Room) RemoveClient(client *Client) {
	r.ClientsMu.Lock()
	defer r.ClientsMu.Unlock()
	delete(r.Clients, client)
}

func (r *Room) Broadcast(sender *Client, mt int, message []byte) {
	r.ClientsMu.Lock()
	defer r.ClientsMu.Unlock()

	for client := range r.Clients {
		if client != sender {
			client.Conn.WriteMessage(mt, message)
		}
	}
}
