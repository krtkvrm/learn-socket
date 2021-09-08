// server.go
package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func serveClapWS(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	client := client{connection: conn, messages: make(chan []byte, 256)}
	hub.addClient <- client
	conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("%d", claps)))

	go client.listenMessage()
	go client.pushMessage()
}

func main() {
	newCore()
	go run()
	http.HandleFunc("/ping", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "pong")
	})
	http.HandleFunc("/claps", func(w http.ResponseWriter, r *http.Request) {
		serveClapWS(w, r)
	})
	http.Handle("/", http.FileServer(http.Dir("./static")))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
