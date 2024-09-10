package relayConn

import (
	"encoding/json"
	"net"
)

// this connects to the relay
func Initialize(relayAddrs string) {
	conn, err := net.Dial("tcp", relayAddrs)
	if err != nil {
		panic(err)
		// handle error
	}

	// we send them to a channel
	rConn = conn
	rReader = json.NewDecoder(conn)

}

func Setup(relayAddrs, idNode string) {
	// start the server
	Initialize(relayAddrs)
	// this is just for controlling other stuff
	go ActualizeSDP()
	go RelayNewConns()
	// i need to wait until i receive the Node ID
	nID = <-IDchan
	// this is only when our node is a client :D
	if idNode != "" {
		ConnectTo(idNode)
		go SendOffering(idNode)

	}

}
