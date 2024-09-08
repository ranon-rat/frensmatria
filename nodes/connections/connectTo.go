package connections

import (
	"fmt"

	"github.com/ranon-rat/frensmatria/core"
	"github.com/ranon-rat/frensmatria/nodes/channels"
)

/*
with this we can handle multiple invitations and information that can come in any moment.
Its quite importan specially for making bigger things
*/
func ConnectToNodes() {
	for {
		// con esto ya puedo mantener mis conexiones y otras cosas
		answerSDP := <-channels.SDPChan
		fmt.Println("new connection")

		switch answerSDP.Kind {
		case core.ConnectTo:
			SDPOfferChan <- answerSDP.SDP
		case core.Confirm:
			SDPAnswerChan <- answerSDP.SDP
		default:
			fmt.Println(answerSDP)
			continue
		}

	}
}
