package main

import (
	"flag"
	"fmt"

	"github.com/ranon-rat/frensmatria/nodes/SDPConn"
	"github.com/ranon-rat/frensmatria/nodes/connections"
	"github.com/ranon-rat/frensmatria/nodes/relayConn"
)

func main() {
	// some simple shit for using it later
	relayAddrs := flag.String("relay", "localhost:8080", "just connect to a relay so we can hole punch")
	idNode := flag.String("node", "", "is just the id that the relay generats, use it to connect with someone else")
	flag.Parse()
	SDPConn.Setup()
	relayConn.Setup(*relayAddrs, *idNode)
	go connections.HandleEventConns()
	fmt.Println("share this ID:", relayConn.GiveID())

	select {}
	// i probably should initialize first the server before anything tbh

}
