package connections

import (
	"github.com/pion/webrtc/v3"
	"github.com/ranon-rat/frensmatria/nodes/src/db"
)

type ConnectionInfo struct {
	CloseChan  chan struct{}
	Connection *webrtc.DataChannel
	MsgChan    chan webrtc.DataChannelMessage
}

type ConnectionID struct {
	Connection *webrtc.DataChannel
	ID         string
}

var (
	ConnInfoChan   = make(chan ConnectionInfo)
	Conns          = make(map[ConnectionID]bool)
	ComparingMap   = []map[string]int{}
	ComparingQs    = []bool{}
	ComparingNodes = 0
	ComparingQ     = false
	CurrentDate    = 0
)

func SetDate() {
	CurrentDate = db.GetLastDate()
	ComparingQ = true
}
