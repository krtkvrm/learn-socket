// server.go
package main

import (
	"fmt"
	"github.com/DataDog/datadog-go/statsd"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

var statsdClient *statsd.Client

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

	statsdClient, _ = statsd.New("0.0.0.0:8125")

	go func() {
		for {
			statsdClient.Count("claps", int64(claps), []string{}, 0)
			statsdClient.Count("clients", int64(clients), []string{}, 0)
			time.Sleep(500 * time.Millisecond)
		}
	}()


	http.HandleFunc("/ping", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "pong")
	})
	http.HandleFunc("/claps", func(w http.ResponseWriter, r *http.Request) {
		serveClapWS(w, r)
	})
	http.Handle("/", http.FileServer(http.Dir("./static")))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
