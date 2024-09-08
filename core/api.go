package core

// le voy a enviar esto ya que este dentro del loop
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
