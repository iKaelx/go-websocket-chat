package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"webSockets/internal/ws"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/ws", ws.HandleWS)

	http.Handle("/", http.FileServer(http.Dir("./frontend")))

	fmt.Println("Server running on http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}