package connections

import (
	"github.com/pion/webrtc/v3"
	"github.com/ranon-rat/frensmatria/nodes/src/core"
)

// we will receive things through the data channel when its opened
func HandleEventConns() {

	for {
		conn := <-ConnInfoChan
		ID := core.RandStringRunes(10)
		cID := ConnectionID{
			ID:         ID,
			Connection: conn,
		}
		Conns[cID] = true
		if ComparingQ {
			IncreaseLifeTime[ID] = make(chan struct{})
			ComparingMap[ID] = make(map[string]int)
			ComparingQs[ID] = true
		}
		go OnOpen(conn, ID)
		conn.OnMessage(func(msg webrtc.DataChannelMessage) {
			OnMessage(cID, msg)
		})
		conn.OnError(func(err error) {
			OnClose(cID)
		})
		conn.OnClose(func() {
			OnClose(cID)
		})
	}
}
