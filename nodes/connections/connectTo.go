package connections

import (
	"fmt"

	"github.com/ranon-rat/frensmatria/core"
	"github.com/ranon-rat/frensmatria/nodes/channels"
)

// okay here is when the hole punching comes
// it takes like a minute to coordinate it will be trying to connect to the other server fast
func ConnectToNodes() {
	for {
		// con esto ya puedo mantener mis conexiones y otras cosas
		answerSDP := <-channels.SDPChan
		fmt.Println("new connection")
		// ya no es necesario que mantenga un canal para poder estar teniendo que estar leyendo los nodos lo cual es
		// bueno
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
