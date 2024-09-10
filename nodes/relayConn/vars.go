package relayConn

import (
	"encoding/json"
	"net"
)

var (
	// relay connection
	rConn net.Conn
	// relay reader
	rReader *json.Decoder
	// node id
	nID string
	// channels for initializing our SDPConn
	ConnChan   = make(chan net.Conn)
	ReaderChan = make(chan *json.Decoder)
	IDchan     = make(chan string)
)

// this is just for setting the variables

// i get the id of the user with
func GiveID() string {

	return nID
}
