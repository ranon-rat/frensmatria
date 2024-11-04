package connections

import (
	"github.com/pion/webrtc/v3"
)

var (
	ConnChan = make(chan *webrtc.DataChannel)
	Conns    = make(map[ConnectionID]bool)

	// its a simple map for mantaining some order
	ComparingMap      = make(map[string]map[string]int)
	ComparingQs       = make(map[string]bool)
	CIncreaseLifeTime = make(map[string]chan struct{})
	Alive             = make(map[string]chan struct{})

	// this is important for keeping some level of structure, in case that its not working correctly i just
	// delete something
	CompareEndChan = make(chan struct{})
	ComparingNodes = 0
	ComparingQ     = false

	LastDate = 0
	// just disconnect it hasnt receive anything in a lot of time
	ExpectedNodes  = 0
	ConnectedNodes = 0

	//msg shit

	MsgCache = make(map[string]bool)
)

// this is for internal usage btw
type ConnectionID struct {
	Connection *webrtc.DataChannel
	ID         string
}
