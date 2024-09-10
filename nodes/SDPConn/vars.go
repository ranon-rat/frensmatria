package SDPConn

import "github.com/pion/webrtc/v3"

var (
	Config = webrtc.Configuration{
		ICEServers: []webrtc.ICEServer{
			{
				URLs: []string{"stun:stun.l.google.com:19302"}, // Servidor STUN
			},
		},
	}
	// they are for communicating between the SDPConn of connectTo and expecting new conns
	SDPAnswerChan = make(chan string)
	SDPOfferChan  = make(chan string)
)
