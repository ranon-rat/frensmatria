package channels

import "github.com/ranon-rat/frensmatria/common"

const (
	Add = "add"
)

var (
	// this is relayConn->SDPConn
	SDPChan = make(chan common.Initial)
	// SDPConn->relayConn
	SDPChanAnswer = make(chan string)
	// SDPConn->relayConn
	SDPChanInivitation = make(chan string)
	// Gematria->connections
	// new     base64json // sending or receiving
	// {content:"example",date:1234567} // check if the content is already in db
	// compare base64json // this is just for sending, or receiving (if you receive this, you shouldnt share it with other nodes)
	// {content:"example",date:12341325}
	// end (this is for the comparing stuf)
	// get dateTime // this is only for getting information
	ConnectionComm = make(chan string)
)
