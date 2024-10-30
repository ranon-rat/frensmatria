package main

import (
	"flag"
	"fmt"

	"github.com/ranon-rat/frensmatria/nodes/src/SDPConn"
	"github.com/ranon-rat/frensmatria/nodes/src/connections"
	"github.com/ranon-rat/frensmatria/nodes/src/relayConn"
	"github.com/ranon-rat/frensmatria/nodes/src/router"
)

/*
so i think that i will be avoiding to add anything quite interesting here
*/

func Setup() {
	// some simple shit for using it later
	relayAddrs := flag.String("relay", "localhost:8080", "just connect to a relay so we can hole punch")
	idNode := flag.String("node", "", "is just the id that the relay generats, use it to connect with someone else")
	flag.Parse()
	// sdp connections
	SDPConn.Setup()
	// relay communication
	relayConn.Setup(*relayAddrs, *idNode)

	// this handles the events
	// okay so this seems to be quite simple
	connections.SetDate()
	go connections.HandleEventConns()
	go connections.SendMessages()

	fmt.Println("share this ID:", relayConn.GiveID())

}

func main() {
	fmt.Println("hello world")
	//fmt.Println(db.GematriaByID(0))
	go Setup()
	router.Setup()
}
