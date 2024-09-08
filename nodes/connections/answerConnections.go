package connections

import (
	"fmt"
	"time"

	"github.com/pion/webrtc/v3"
	"github.com/ranon-rat/frensmatria/nodes/channels"
)

func ONAnswer() {
	for {
		peerConn, err := webrtc.NewPeerConnection(Config)
		if err != nil {
			panic(err)
		}

		// we create a channel of data
		peerConn.OnDataChannel(func(d *webrtc.DataChannel) {
			fmt.Println("we are open")
			d.OnOpen(func() {
				fmt.Println("finally")

				for {
					if d.SendText("sup dude, message sended from this client") != nil {
						break
					}

					time.Sleep(time.Second)

				}
			})

			d.OnMessage(func(msg webrtc.DataChannelMessage) {
				fmt.Printf("message: %s\n", string(msg.Data))
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
		fmt.Println("everything has passed")
		channels.SDPChanAnswer <- peerConn.LocalDescription().SDP
	}
}
