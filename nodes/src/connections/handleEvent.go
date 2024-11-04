package connections

import (
	"github.com/pion/webrtc/v3"
	"github.com/ranon-rat/frensmatria/nodes/src/core"
)

// we will receive things through the data channel when its opened
func HandleEventConns() {

	for {
		conn := <-ConnChan
		ID := core.RandStringRunes(10)
		cID := ConnectionID{
			ID:         ID,
			Connection: conn,
		}
		Conns[cID] = true
		Alive[ID] = make(chan struct{})
		if ComparingQ {
			CIncreaseLifeTime[ID] = make(chan struct{})
			ComparingMap[ID] = make(map[string]int)
			ComparingQs[ID] = true
		}
		go OnOpen(cID)
		go conn.OnClose(func() {
			OnClose(cID)
		})
		go conn.OnMessage(func(msg webrtc.DataChannelMessage) {
			OnMessage(cID, msg)
		})
		go conn.OnError(func(err error) {
			OnClose(cID)
		})

	}
}
