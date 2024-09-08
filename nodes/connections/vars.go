package connections

import "github.com/pion/webrtc/v3"

var (
	Config = webrtc.Configuration{
		ICEServers: []webrtc.ICEServer{
			{
				URLs: []string{"stun:stun.l.google.com:19302"}, // Servidor STUN
			},
		},
	}
	// so these are the escential variables hmm

	SDPAnswerChan = make(chan string)
	SDPOfferChan  = make(chan string)
)
