package main

import (
	"flag"
	"fmt"

	"github.com/fatih/color"
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

}

func main() {
	fmt.Println(color.New(color.Bold).AddRGB(100, 0, 255).SprintFunc()(
		`
┏┓               •  
┣ ┏┓┏┓┏┓┏┏┳┓┏┓╋┏┓┓┏┓
┻ ┛ ┗ ┛┗┛┛┗┗┗┻┗┛ ┗┗┻
created by @tecnopsychosis(AQ 333)

`)) // tmplr ascii art on https://patorjk.com/

	relayAddrs := flag.String("relay", "localhost:9090", "just connect to a relay so we can hole punch")
	idNode := flag.String("node", "", "is just the id that the relay generats, use it to connect with someone else")
	port := flag.String("port", "6969", "its the port for the local server")
	update := flag.Bool("not-update", false, "its for updating the db once the service starts")
	flag.Parse()

	Setup(
		*relayAddrs, *idNode, !*update)
	c := color.New(color.Bold).AddRGB(0, 255, 0).SprintFunc()
	fmt.Printf("%s %s \n\n", c("share this ID:"), relayConn.GiveID())

	connections.Setup(!*update)

	router.Setup(*port)
}
