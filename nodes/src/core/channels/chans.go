package channels

// this is for making it easy to communicate between  modules
type Initial struct {
	Kind int    `json:"kind"`
	SDP  string `json:"direction"`
}

// there
type IDResponse struct {
	ID string `json:"id"`
}

type WantConnect struct {
	IDNode   string `json:"idNode"`
	SDPOffer string `json:"SDP"` // if this is empty, that means that i shouldnt send it to the IDNode
}

var (
	// this is relayConn->SDPConn
	SDPChan = make(chan Initial)
	// SDPConn->relayConn
	SDPChanAnswer = make(chan string)
	// SDPConn->relayConn
	SDPChanInivitation = make(chan string)
)
