package main

import (
	"flag"
	"fmt"
	"math/rand"
	"strconv"

	"github.com/fatih/color"
	"github.com/ranon-rat/frensmatria/nodes/src/SDPConn"
	"github.com/ranon-rat/frensmatria/nodes/src/connections"
	"github.com/ranon-rat/frensmatria/nodes/src/core"
	"github.com/ranon-rat/frensmatria/nodes/src/relayConn"
	"github.com/ranon-rat/frensmatria/nodes/src/router"
)

func Setup(relayAddrs, idNode, password string, update bool) {
	// sdp connections
	SDPConn.Setup()
	// relay communication
	relayConn.Setup(relayAddrs, idNode, password)
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
	port := flag.Int("http-server", 0, "its the port for the local server")
	update := flag.Bool("update", false, "its for updating the db once the service starts")
	password := flag.String("password", "", "password for connecting with the relay")
	username := flag.String("username", "anonymous"+strconv.Itoa(rand.Intn(100)), "")
	flag.Parse()
	core.SetUsername(*username)
	Setup(
		*relayAddrs, *idNode, *password, *update)
	c := color.New(color.Bold).AddRGB(0, 255, 0).SprintFunc()
	fmt.Printf("%s %s \n", c("share this ID:"), relayConn.GiveID())
	fmt.Printf("%s %s \n\n", c("username:"), core.Username)

	connections.Setup(*update)
	if *port != 0 {
		go router.Setup(strconv.Itoa(*port))
	}
	select {}
}
