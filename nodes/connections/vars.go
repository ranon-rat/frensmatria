package connections

import "github.com/pion/webrtc/v3"

type ConnectionInfo struct {
	CloseChan  chan struct{}
	Connection *webrtc.DataChannel
	MsgChan    chan webrtc.DataChannelMessage
}

var (
	ConnInfoChan = make(chan ConnectionInfo)
	Conns        = make(map[*webrtc.DataChannel]bool)
)
