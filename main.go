package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func websocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}

		// received and show message
		log.Printf("Received message: %s\n", message)

		msgReceived := fmt.Sprintf("Message received: %s\n", message)

		// Send the message back to the client
		err = conn.WriteMessage(websocket.TextMessage, []byte(msgReceived))
		if err != nil {
			log.Println(err)
			return
		}
	}
}

func main() {
	http.HandleFunc("/ws-message", websocketHandler)
	log.Fatal(http.ListenAndServe(":2323", nil))
}
