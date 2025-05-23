package Websocket

import (
	"encoding/json"
	"fmt"
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

// Instance is the global manager instance
var Instance *Manager

func init() {
	Instance = NewManager()
}

// FindClientByID finds a client in a specific room
func (m *Manager) FindClientByID(roomID string, clientID string) *Client {
	m.RoomsMu.RLock()
	defer m.RoomsMu.RUnlock()

	room := m.Rooms[roomID]
	if room == nil {
		return nil
	}

	room.ClientsMu.Lock()
	defer room.ClientsMu.Unlock()

	for client := range room.Clients {
		if client.ID == clientID {
			return client
		}
	}
	return nil
}

// SendToClient sends a message to a specific client in a specific room
func (m *Manager) SendToClient(roomID string, clientID string, messageType string, data interface{}) error {
	client := m.FindClientByID(roomID, clientID)
	if client == nil {
		return fmt.Errorf("client not found in room %s", roomID)
	}

	payload := map[string]interface{}{
		"type": messageType,
		"data": data,
	}

	msg, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	return client.Conn.WriteMessage(1, msg) // 1 for text message
}

// RoomStats holds information about a room
type RoomStats struct {
	ClientCount int      `json:"client_count"`
	RoomID      string   `json:"room_id"`
	ClientIDs   []string `json:"client_ids"`
}

// GetRoomStats returns statistics about a specific room
func (m *Manager) GetRoomStats(roomID string) (*RoomStats, error) {
	m.RoomsMu.RLock()
	defer m.RoomsMu.RUnlock()

	room := m.Rooms[roomID]
	if room == nil {
		return nil, fmt.Errorf("room %s not found", roomID)
	}

	room.ClientsMu.Lock()
	defer room.ClientsMu.Unlock()

	clientIDs := make([]string, 0, len(room.Clients))
	for client := range room.Clients {
		clientIDs = append(clientIDs, client.ID)
	}

	return &RoomStats{
		ClientCount: len(room.Clients),
		RoomID:      roomID,
		ClientIDs:   clientIDs,
	}, nil
}

// BroadcastToRoom sends a message to all clients in a room
func (m *Manager) BroadcastToRoom(roomID string, messageType string, data interface{}) error {
	room := m.GetRoom(roomID)
	if room == nil {
		return fmt.Errorf("room %s not found", roomID)
	}

	payload := map[string]interface{}{
		"type": messageType,
		"data": data,
	}

	msg, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	room.ClientsMu.Lock()
	defer room.ClientsMu.Unlock()

	for client := range room.Clients {
		if err := client.Conn.WriteMessage(1, msg); err != nil {
			// Log error but continue broadcasting to other clients
			fmt.Printf("Error broadcasting to client %s: %v\n", client.ID, err)
		}
	}
	return nil
}
