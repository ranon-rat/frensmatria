package SDPConn

import (
	"fmt"

	"github.com/pion/webrtc/v3"
	"github.com/ranon-rat/frensmatria/nodes/src/core/channels"
)

func EnterConn() {
	for {
		peerConn, err := webrtc.NewPeerConnection(Config)
		if err != nil {
			panic(err)
		}

		// we create a channel of data
		peerConn.OnDataChannel(func(d *webrtc.DataChannel) {
			dcHandler(d)
		})
		SDP := <-SDPOfferChan
		offer := webrtc.SessionDescription{
			Type: webrtc.SDPTypeOffer,
			SDP:  SDP,
		}

		if err = peerConn.SetRemoteDescription(offer); err != nil {
			fmt.Println(err)
			continue
		}

		// Crear una respuesta SDP
		answer, err := peerConn.CreateAnswer(nil)
		if err != nil {
			fmt.Println(err)
			continue

		}
		GetICE(peerConn, answer)

		channels.SDPChanAnswer <- peerConn.LocalDescription().SDP
	}
}
