package SDPConn

import (
	"fmt"

	"github.com/pion/webrtc/v3"
	"github.com/ranon-rat/frensmatria/nodes/channels"
	"github.com/ranon-rat/frensmatria/nodes/connections"
)

func ONAnswer() {
	for {
		peerConn, err := webrtc.NewPeerConnection(Config)
		if err != nil {
			panic(err)
		}

		// we create a channel of data
		peerConn.OnDataChannel(func(d *webrtc.DataChannel) {
			closeChan := make(chan struct{})
			msgChan := make(chan webrtc.DataChannelMessage)
			d.OnOpen(func() {
				connections.ConnInfoChan <- connections.ConnectionInfo{
					CloseChan:  closeChan,
					MsgChan:    msgChan,
					Connection: d,
				}

			})
			d.OnClose(func() {
				closeChan <- struct{}{}

			})
			d.OnMessage(func(msg webrtc.DataChannelMessage) {
				msgChan <- msg
			})
		})
		SDP := <-SDPOfferChan
		offer := webrtc.SessionDescription{
			Type: webrtc.SDPTypeOffer,
			SDP:  SDP,
		}
		err = peerConn.SetRemoteDescription(offer)
		if err != nil {
			fmt.Println("100", err)

			continue
		}

		// Crear una respuesta SDP
		answer, err := peerConn.CreateAnswer(nil)
		if err != nil {
			fmt.Println("108", err)
			continue

		}
		gatherComplete := webrtc.GatheringCompletePromise(peerConn)

		err = peerConn.SetLocalDescription(answer)
		if err != nil {
			fmt.Println("116", err)
			continue
		}

		<-gatherComplete
		channels.SDPChanAnswer <- peerConn.LocalDescription().SDP
	}
}
