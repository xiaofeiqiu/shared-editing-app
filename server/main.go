package main

import (
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

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan []byte)

// handleConnections handles incoming WebSocket connections
func handleConnections(w http.ResponseWriter, r *http.Request) {
    log.Println("WebSocket connection attempt")
    ws, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Printf("Error upgrading to WebSocket: %v", err)
        return
    }
    defer ws.Close()

    log.Println("WebSocket connection established")
    clients[ws] = true

    for {
        messageType, message, err := ws.ReadMessage()
        if err != nil {
            log.Printf("Error reading message: %v", err)
            delete(clients, ws)
            break
        }

        if messageType == websocket.TextMessage {
            log.Printf("Received text message: %s", message)
        } else if messageType == websocket.BinaryMessage {
            log.Println("Received binary message")
        }

        broadcast <- message
    }
}

// handleMessages broadcasts messages to all connected clients
func handleMessages() {
    for {
        msg := <-broadcast
        for client := range clients {
            err := client.WriteMessage(websocket.BinaryMessage, msg)
            if err != nil {
                log.Printf("Error writing message: %v", err)
                client.Close()
                delete(clients, client)
            } else {
                log.Printf("Message sent to client: %v, message: %s", client.RemoteAddr(), string(msg))
            }
        }
    }
}

func main() {
    // Serve the index.html file
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "index.html")
    })

    http.HandleFunc("/ws/quill-demo", handleConnections)
    go handleMessages()

    log.Println("Server started on :8080")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatalf("ListenAndServe: %v", err)
    }
}
