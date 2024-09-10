package channels

import (
	"github.com/ranon-rat/frensmatria/core"
)

// this is for making it easy to communicate between  modules

var (
	// this is relayConn->SDPConn
	SDPChan = make(chan core.Initial)
	// SDPConn->relayConn
	SDPChanAnswer = make(chan string)
	// SDPConn->relayConn
	SDPChanInivitation = make(chan string)
)
