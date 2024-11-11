package controllers

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/ranon-rat/frensmatria/nodes/src/core"
	"github.com/ranon-rat/frensmatria/nodes/src/core/channels"
)

func SetupWebsocket(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrade.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer ws.Close()

	// some stuff for do this work
	// lo que estoy haciendo aqui es checar si el canal que se a ingresado , en el caso de que no exista , se cierra

	clients[ws] = true

	go SendMessages() // send the messages
	ReceiveMessages(ws)

}

func SendMessages() {
	for {
		msg := <-Message // wait for the message

		for client := range clients {

			if err := client.WriteJSON(msg); err != nil { // if the socket is closed

				delete(clients, client)

				// if no one is in the channel,the channel and the message is deleted

				return
			}
		}
	}

}

func ReceiveMessages(ws *websocket.Conn) {
	for { // receive the messsages
		var msg core.Messages

		// parse the message
		if err := ws.ReadJSON(&msg); err != nil { // if the websocket is closed
			// if the websocket is closed
			log.Println(err)
			delete(clients, ws)
			// if no one is in the channel,the channel and the message is deleted

			return

		}
		msg.Timestamp = int(time.Now().UnixNano())
		msg.Username = core.Username
		go channels.SendMessage(fmt.Sprintf("message %s", core.Object2Base64(msg)), "")
		Message <- msg
	}

}
