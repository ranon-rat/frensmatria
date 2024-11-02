package main

import (
	"flag"
	"fmt"

	"github.com/ranon-rat/frensmatria/nodes/src/SDPConn"
	"github.com/ranon-rat/frensmatria/nodes/src/connections"
	"github.com/ranon-rat/frensmatria/nodes/src/relayConn"
	"github.com/ranon-rat/frensmatria/nodes/src/router"
)

func Setup(relayAddrs, idNode string, update bool) {
	// sdp connections
	SDPConn.Setup()
	// relay communication
	relayConn.Setup(relayAddrs, idNode)
	// this handles the events
	connections.Setup(update)
	fmt.Println("share this ID:", relayConn.GiveID())
}

func main() {

	relayAddrs := flag.String("relay", "localhost:9090", "just connect to a relay so we can hole punch")
	idNode := flag.String("node", "", "is just the id that the relay generats, use it to connect with someone else")
	port := flag.String("port", "6969", "its the port for the local server")
	update := flag.Bool("not-update", false, "its for updating the db once the service starts")
	flag.Parse()

	go Setup(
		*relayAddrs, *idNode, !*update)
	router.Setup(*port)
}
