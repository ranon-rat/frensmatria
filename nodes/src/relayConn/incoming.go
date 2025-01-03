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
		for _, id := range res.NodesID {
			go ConnectTo(id)

		}
		IDchan <- res.ID
		channels.HowManyNodes <- len(res.NodesID)

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
			Password: Password,
		})
		// with this i just manage other stuff
	}
}
