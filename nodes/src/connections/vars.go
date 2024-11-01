package connections

import (
	"math/rand"

	"github.com/pion/webrtc/v3"
	"github.com/ranon-rat/frensmatria/nodes/src/db"
)

var (
	ConnInfoChan = make(chan ConnectionInfo)
	Conns        = make(map[ConnectionID]bool)

	// its a simple map for mantaining some order
	ComparingMap     = make(map[string]map[string]int)
	ComparingQs      = make(map[string]bool)
	IncreaseLifeTime = make(map[string]chan struct{})

	// this is important for keeping some level of structure, in case that its not working correctly i just
	// delete something
	CompareEndChan = make(chan struct{})
	ComparingNodes = 0
	ComparingQ     = false

	LastDate    = 0
	letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

func SetDate() {
	LastDate = db.GetLastDate()
	ComparingQ = true
}

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

type ConnectionInfo struct {
	CloseChan  chan struct{}
	Connection *webrtc.DataChannel
	MsgChan    chan webrtc.DataChannelMessage
}

type ConnectionID struct {
	Connection *webrtc.DataChannel
	ID         string
}
