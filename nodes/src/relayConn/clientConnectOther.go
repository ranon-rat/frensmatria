package relayConn

import (
	"encoding/json"

	"github.com/ranon-rat/frensmatria/common"
	"github.com/ranon-rat/frensmatria/nodes/src/core/channels"
)

func ConnectTo(id string) {
	IDConnectChan <- id
}
func SendOffering() {
	for {
		id := <-IDConnectChan
		json.NewEncoder(rConn).Encode(common.WantConnect{
			IDNode: id,
		})

		SDP := <-channels.SDPChanAnswer
		json.NewEncoder(rConn).Encode(common.WantConnect{
			SDPOffer: SDP,
			IDNode:   id,
		})
	}
}
