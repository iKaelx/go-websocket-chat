package ws

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"webSockets/internal/ai"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func HandleWS(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}

		reply := ai.GetGroqResponse(string(msg))

		conn.WriteMessage(websocket.TextMessage, []byte(reply))
	}
}