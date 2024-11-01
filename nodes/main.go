package main

import (
	"flag"
	"fmt"

	"github.com/ranon-rat/frensmatria/nodes/src/SDPConn"
	"github.com/ranon-rat/frensmatria/nodes/src/connections"
	"github.com/ranon-rat/frensmatria/nodes/src/relayConn"
	"github.com/ranon-rat/frensmatria/nodes/src/router"
)

func Setup(relayAddrs, idNode string) {
	// some simple shit for using it later

	// sdp connections
	SDPConn.Setup()
	// relay communication
	relayConn.Setup(relayAddrs, idNode)

	// this handles the events
	// okay so this seems to be quite simple
	go connections.Setup()
	fmt.Println("share this ID:", relayConn.GiveID())
}

func main() {

	relayAddrs := flag.String("relay", "localhost:8080", "just connect to a relay so we can hole punch")
	idNode := flag.String("node", "", "is just the id that the relay generats, use it to connect with someone else")
	port := flag.String("port", "8080", "its the port for the local server")
	flag.Parse()

	go Setup(
		*relayAddrs, *idNode)
	router.Setup(*port)
}
