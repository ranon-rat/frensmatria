package relayConn

import (
	"encoding/json"

	"github.com/ranon-rat/frensmatria/common"
	"github.com/ranon-rat/frensmatria/nodes/src/core/channels"
)

/*
with this we can handle the relay answers
*/
func RelayNewConns() {
	defer rConn.Close()
	if nID == "" {
		var res common.IDResponse
		// i continue iin my journey
		if rReader.Decode(&res) != nil {
			panic("relay problems")
		}
		IDchan <- res.ID
	}
	for {
		var body common.Initial
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

/*
with this we update the SDP, its quite important and we have to take it into our mind
*/
func ActualizeSDP() {

	for {
		SDP := <-channels.SDPChanInivitation
		json.NewEncoder(rConn).Encode(common.WantConnect{
			SDPOffer: SDP,
			IDNode:   "",
		})
		// with this i just manage other stuff
	}
}
