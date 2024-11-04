package common

type Initial struct {
	Kind int    `json:"kind"`
	SDP  string `json:"direction"`
}

// there
type IDResponse struct {
	ID      string   `json:"id"`
	NodesID []string `json:"nodes"`
}

type WantConnect struct {
	Password string `json:"password"`
	IDNode   string `json:"idNode"`
	SDPOffer string `json:"SDP"` // if this is empty, that means that i shouldnt send it to the IDNode
}
