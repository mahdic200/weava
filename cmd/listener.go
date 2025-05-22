package cmd

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/coder/websocket"
	"github.com/spf13/cobra"
)

// listenerCmd represents the listener command
var listenerCmd = &cobra.Command{
	Use:   "listener",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		conn, _, err := websocket.Dial(ctx, "ws://localhost:8080", nil)
		if err != nil {
			log.Fatalf("Failed to connect to server: %v", err)
		}
		defer conn.Close(websocket.StatusNormalClosure, "Client closing connection")
		fmt.Println("Connected to the echo server. Type messages to send (Ctrl+C to exit):")
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			message := scanner.Text()

			// Send message to the server
			err = conn.Write(ctx, websocket.MessageText, []byte(message))
			if err != nil {
				log.Printf("Writer error: %v", err)
				return
			}

			// Read echoed message from the server
			_, data, err := conn.Read(ctx)
			if err != nil {
				log.Printf("Read error: %v", err)
				return
			}

			fmt.Printf("Echoed: %s\n", data)
		}
	},
}

func init() {
	rootCmd.AddCommand(listenerCmd)
}
