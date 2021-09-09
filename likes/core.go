package main

import (
	"bytes"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"strings"
	"time"
)

var claps = 0
var clients = 0

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

type client struct {
	connection *websocket.Conn
	messages   chan []byte
}

type Hub struct {
	clients      map[client]bool
	broadcast    chan []byte
	addClient    chan client
	removeClient chan client
}

var hub Hub

func newCore() {
	hub = Hub{
		clients:      make(map[client]bool),
		addClient:    make(chan client),
		removeClient: make(chan client),
		broadcast:    make(chan []byte),
	}
}

func run() {
	for {
		select {
		case client := <-hub.addClient:
			hub.clients[client] = true
			clients++
			fmt.Printf("Clients = %d\n", clients)
		case client := <-hub.removeClient:
			if _, ok := hub.clients[client]; ok {
				delete(hub.clients, client)
				close(client.messages)
			}
			clients--
			fmt.Printf("Clients = %d\n", clients)
		case message := <-hub.broadcast:
			for client := range hub.clients {
				select {
				case client.messages <- message:
				default:
					close(client.messages)
					delete(hub.clients, client)
				}
			}
		}
	}
}

func (c client) listenMessage() {
	defer func() {
		hub.removeClient <- c
		c.connection.Close()
	}()
	for {
		_, message, err := c.connection.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		if strings.TrimSpace(string(message)) == "c" {
			claps++
		}
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		hub.broadcast <- []byte(fmt.Sprintf("%d", claps))
	}
}

func (c client) pushMessage() {
	ticker := time.NewTicker(1 * time.Second)
	defer func() {
		ticker.Stop()
		c.connection.Close()
	}()

	for {
		select {
		case message, ok := <-c.messages:
			c.connection.SetWriteDeadline(time.Now().Add(10 * time.Second))
			if !ok {
				c.connection.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.connection.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.connection.SetWriteDeadline(time.Now().Add(10 * time.Second))
			if err := c.connection.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}
