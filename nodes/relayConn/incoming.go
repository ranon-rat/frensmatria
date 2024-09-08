package relayConn

import (
	"encoding/json"

	"github.com/ranon-rat/frensmatria/core"
	"github.com/ranon-rat/frensmatria/nodes/channels"
)

func RelayNewConns() {
	defer rConn.Close()
	if nID == "" {
		// with this i just manage other stuff
		var res core.IDResponse
		// i continue iin my journey
		if rReader.Decode(&res) != nil {
			panic("relay problems")
		}
		// we send them to a channel
		IDchan <- res.ID
	}
	for {
		var body core.Initial
		if rReader.Decode(&body) != nil {
			break
		}
		if body.SDP == "" {
			// i cannot send this
			continue
		}

		channels.SDPChan <- body
	}

}
func ActualizeSDP() {

	for {
		SDP := <-channels.SDPChanInivitation
		json.NewEncoder(rConn).Encode(core.WantConnect{
			SDPOffer: SDP,
			IDNode:   "",
		})

		// with this i just manage other stuff

	}
}
