package SDPConn

import (
	"github.com/pion/webrtc/v3"
	"github.com/ranon-rat/frensmatria/nodes/src/core/channels"
)

/*
this function works for mantaining multiple client SDPConn between nodes
its quite useful specially for mantaining some kind of stability

use it in a goroutine
*/
func IncomingConn() string {
	// this is for handling multiple SDPConn
	for {

		peerConn, err := webrtc.NewPeerConnection(Config)
		if err != nil {
			panic(err)
		}
		//the panic stuff is just in case something is not workign
		dataChannel, err := peerConn.CreateDataChannel("data", nil)
		if err != nil {
			panic(err)
		}

		// so we add some events here for talkign with the other folders
		dcHandler(dataChannel)

		// we create an SDP offer
		offer, err := peerConn.CreateOffer(nil)
		if err != nil {
			panic(err)
		}
		GetICE(peerConn, offer)

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
