package HomeController

import (
	"log"

	"github.com/gofiber/contrib/websocket"
	"github.com/mahdic200/weava/Websocket"
	"github.com/mahdic200/weava/Utils"
)

var manager = Websocket.NewManager()

var Index = websocket.New(func(c *websocket.Conn) {
	// Get client ID from URL parameters
	clientID := Utils.StandardRandomString(16)
	roomID := "chat" // Or get from params/query

	room := manager.GetRoom(roomID)
	if room == nil {
		room = manager.CreateRoom(roomID)
	}

	client := &Websocket.Client{
		ID:   clientID,
		Conn: c,
	}

	room.AddClient(client)
	defer room.RemoveClient(client)

	// Message handling loop
	var (
		mt  int    // Message type
		msg []byte // Message content
		err error  // Error handling
	)

	// Infinite loop to handle messages
	for {
		// Read incoming message
		if mt, msg, err = c.ReadMessage(); err != nil {
			log.Println("read:", err)
			break // Break loop if error occurs
		}
		log.Printf("recv from %s: %s", clientID, msg)

		// Send message to all other clients
		room.Broadcast(client, mt, msg)
	}
})
