package SDPConn

import "github.com/pion/webrtc/v3"

// this works for getting the full SDP, is important for connecting with other nodes

func GetICE(peerConn *webrtc.PeerConnection, desc webrtc.SessionDescription) {
	gatherComplete := webrtc.GatheringCompletePromise(peerConn)

	// We put it in our  local description
	err := peerConn.SetLocalDescription(desc)
	if err != nil {
		panic(err)
	}

	// we wait until we gather all the data
	<-gatherComplete
}
