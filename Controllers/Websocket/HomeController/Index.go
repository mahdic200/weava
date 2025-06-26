package HomeController

import (
	"fmt"
	"log"

	"github.com/gofiber/contrib/websocket"
	"github.com/mahdic200/weava/Models"
	"github.com/mahdic200/weava/Websocket"
)

var manager = Websocket.NewManager()

var Index = websocket.New(func(c *websocket.Conn) {
	/* Getting current logged in client's ID from user object which is
	initiated from AuthMiddleware, I skipped the error check because
	AuthMiddleware makes sure that we have a logged in user, and if we
	don't pass the AuthMiddleware accidentally to the router or we use
	the wrong keyword for Locals method the program will panic */
	var user, _ = c.Locals("user").(Models.User)
	clientID := fmt.Sprintf("%v", user.Id)
	roomID := "chat"

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
