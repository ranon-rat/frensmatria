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
	ConnChan <- conn
}

// this is for starting the SDPConn and everythign
func SetupVariables() {

	rConn = <-ConnChan
	rReader = json.NewDecoder(rConn)
}
func SetupID() {
	nID = <-IDchan

}

func Setup(relayAddrs, idNode string) {
	go Initialize(relayAddrs)
	SetupVariables()
	go ActualizeSDP()
	go RelayNewConns()
	SetupID()
	if idNode != "" {
		ConnectTo(idNode)
		go SendOffering(idNode)

	}

}
