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
	// i sent him the SDP

	ConnChan <- conn
}

// this is for starting the connections and everythign
func SetupVariables() {

	rConn = <-ConnChan
	rReader = json.NewDecoder(rConn)

}
func SetupID() {
	nID = <-IDchan

}
