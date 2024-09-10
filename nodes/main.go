package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/ranon-rat/frensmatria/nodes/SDPConn"
	"github.com/ranon-rat/frensmatria/nodes/connections"
	"github.com/ranon-rat/frensmatria/nodes/relayConn"
)

func main() {
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
	go connections.HandleEventConns()
	go connections.SendMessages()

	fmt.Println("share this ID:", relayConn.GiveID())
	go func() {
		for {
			fmt.Print("> ")
			reader := bufio.NewReader(os.Stdin)
			content, _ := reader.ReadString('\n')
			connections.SetMSG(content, "")
		}
	}()
	select {}
}
