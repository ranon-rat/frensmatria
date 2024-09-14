package connections

import "github.com/pion/webrtc/v3"

type ConnectionInfo struct {
	CloseChan  chan struct{}
	Connection *webrtc.DataChannel
	MsgChan    chan webrtc.DataChannelMessage
}
type Message struct {
	ID      string
	Content string
}
type ConnectionID struct {
	Connection *webrtc.DataChannel
	ID         string
}

var (
	ConnInfoChan = make(chan ConnectionInfo)
	Conns        = make(map[ConnectionID]bool)
	MsgChan      = make(chan Message)
)
