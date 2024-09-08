package relayConn

import (
	"encoding/json"
	"fmt"

	"github.com/ranon-rat/frensmatria/core"
	"github.com/ranon-rat/frensmatria/nodes/channels"
)

func ConnectTo(id string) {
	json.NewEncoder(rConn).Encode(core.WantConnect{
		IDNode: id,
	})
}
func SendOffering(id string) {
	for {
		fmt.Println("sending sdp")
		SDP := <-channels.SDPChanAnswer
		fmt.Println("sending sdp")
		json.NewEncoder(rConn).Encode(core.WantConnect{
			SDPOffer: SDP,
			IDNode:   id,
		})
	}
}
