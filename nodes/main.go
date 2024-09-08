package main

import (
	"flag"
	"fmt"

	"github.com/ranon-rat/frensmatria/nodes/connections"
	"github.com/ranon-rat/frensmatria/nodes/relayConn"
)

func main() {
	// some simple shit for using it later
	relayAddrs := flag.String("relay", "localhost:8080", "just connect to a relay so we can hole punch")
	idNode := flag.String("node", "", "is just the id that the relay generats, use it to connect with someone else")
	flag.Parse()

	go connections.StartPeer()
	go connections.ONAnswer()
	go relayConn.Initialize(*relayAddrs)
	go relayConn.ActualizeSDP()

	relayConn.SetupVariables()
	go relayConn.RelayNewConns()
	relayConn.SetupID()
	fmt.Println("share this ID:", relayConn.GiveID())

	if *idNode != "" {
		fmt.Println(*idNode)
		relayConn.ConnectTo(*idNode)
		go relayConn.SendOffering(*idNode)

	}

	go connections.ConnectToNodes()

	select {}
	// i probably should initialize first the server before anything tbh

}
