package connections

import (
	"fmt"
	"time"

	"github.com/pion/webrtc/v3"
	"github.com/ranon-rat/frensmatria/nodes/channels"
)

/*
this function works for mantaining multiple client connections between nodes
its quite useful specially for mantaining some kind of stability

use it in a goroutine
*/
func OfferConnections() string {
	// this is for handling multiple connections
	for {

		peerConn, err := webrtc.NewPeerConnection(Config)
		if err != nil {
			panic(err)
		}

		dataChannel, err := peerConn.CreateDataChannel("data", nil)
		if err != nil {
			panic(err)
		}

		// so this is just when it opens
		dataChannel.OnOpen(func() {
			fmt.Println("we have our channel open sheesh ")
			for {

				err := dataChannel.SendText("sup dude")
				if err != nil {
					break
				}
				time.Sleep(time.Second)
			}
		})

		// this will handle the messages received
		dataChannel.OnMessage(func(msg webrtc.DataChannelMessage) {
			fmt.Printf("new message: %s\n", string(msg.Data))
		})

		// we create an SDP offer
		offer, err := peerConn.CreateOffer(nil)
		if err != nil {
			panic(err)
		}
		gatherComplete := webrtc.GatheringCompletePromise(peerConn)

		// We put it in our  local description
		err = peerConn.SetLocalDescription(offer)
		if err != nil {
			panic(err)
		}

		// we wait until we gather all the data
		<-gatherComplete
		// i share this with other possible nodes, this is just for updating the sdp
		channels.SDPChanInivitation <- peerConn.LocalDescription().SDP
		// i wait until someone wants to join
		answerSDP := <-SDPAnswerChan

		answer := webrtc.SessionDescription{
			Type: webrtc.SDPTypeAnswer,
			SDP:  answerSDP,
		}

		if peerConn.SetRemoteDescription(answer) != nil {
			continue
		}

	}
}
