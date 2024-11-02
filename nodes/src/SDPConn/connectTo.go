package SDPConn

import (
	"github.com/ranon-rat/frensmatria/nodes/src/core"
	"github.com/ranon-rat/frensmatria/nodes/src/core/channels"
)

/*
with this we can handle multiple invitations and information that can come in any moment.
Its quite important specially for making bigger things
*/
func ConnectToNodes() {
	for {

		answerSDP := <-channels.SDPChan
		switch answerSDP.Kind {
		case core.ConnectTo:
			SDPOfferChan <- answerSDP.SDP
		case core.Confirm:
			SDPAnswerChan <- answerSDP.SDP
		default:
			continue
		}

	}
}
