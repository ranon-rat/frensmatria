package channels

import "github.com/ranon-rat/frensmatria/core"

// this is for making it easy to communicate between  modules

var (
	SDPChan            = make(chan core.Initial)
	SDPChanAnswer      = make(chan string)
	SDPChanInivitation = make(chan string)
)
