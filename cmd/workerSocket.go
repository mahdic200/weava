/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/coder/websocket"
	"github.com/spf13/cobra"
)

// workerSocketCmd represents the workerSocket command
var workerSocketCmd = &cobra.Command{
	Use:   "workerSocket",
	Short: "Starts the websocket for failed tasks worker",
	Long:  `This command initializes the failed tasks worker websocket to handle failed jobs .`,
	Run: func(cmd *cobra.Command, args []string) {

		http.HandleFunc("/", echoHandler)
		fmt.Println("Echo server started at ws://localhost:8080")
		log.Fatal(http.ListenAndServe(":8080", nil))
	},
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := websocket.Accept(w, r, &websocket.AcceptOptions{
		InsecureSkipVerify: true,
	})

	if err != nil {
		log.Println("Failed to accept connection:", err)
		return
	}
	defer conn.Close(websocket.StatusNormalClosure, "")

	for {
		msgType, data, err := conn.Read(r.Context())
		if websocket.CloseStatus(err) != -1 {
			log.Println("Connection closed:", err)
			return
		}
		if err != nil {
			log.Println("Read error:", err)
			return
		}
		var parsedData map[string]any
		if err := json.Unmarshal([]byte(data), &parsedData); err != nil {
			log.Printf("Could not parse data")
		} else {
			for key, value := range parsedData {
				fmt.Printf("%v, %v\n", key, value)
			}
		}
		log.Printf("Received: %s\n", parsedData)
		err = conn.Write(r.Context(), msgType, data)
		if err != nil {
			log.Println("Writer error:", err)
			return
		}
	}
}

func init() {
	rootCmd.AddCommand(workerSocketCmd)
}
