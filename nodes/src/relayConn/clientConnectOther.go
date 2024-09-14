package relayConn

import (
	"encoding/json"

	"github.com/ranon-rat/frensmatria/core"
	"github.com/ranon-rat/frensmatria/nodes/src/core/channels"
)

func ConnectTo(id string) {
	json.NewEncoder(rConn).Encode(core.WantConnect{
		IDNode: id,
	})
}
func SendOffering(id string) {
	for {
		SDP := <-channels.SDPChanAnswer
		json.NewEncoder(rConn).Encode(core.WantConnect{
			SDPOffer: SDP,
			IDNode:   id,
		})
	}
}
