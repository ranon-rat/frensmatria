package relayConn

import (
	"encoding/json"
	"net"
)

func Initialize(relayAddrs string) {
	conn, err := net.Dial("tcp", relayAddrs)
	if err != nil {
		panic(err)
		// handle error
	}
	// i sent him the SDP

	ConnChan <- conn
}

func SetupVariables() {

	rConn = <-ConnChan
	rReader = json.NewDecoder(rConn)

}
func SetupID() {
	nID = <-IDchan

}
