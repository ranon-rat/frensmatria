package connections

import "github.com/pion/webrtc/v3"

// we will receive things through the data channel
func HandleEventConns() {

	for {
		conn := <-ConnInfoChan
		ID := RandStringRunes(10)
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
		// so this will be listening when we close the channel
		go OnOpen(conn, ID)
		conn.OnMessage(func(msg webrtc.DataChannelMessage) {
			OnMessage(conn, msg, ID)
		})
		conn.OnError(func(err error) {
			OnClose(cID)
		})
		conn.OnClose(func() {
			OnClose(cID)
		})
	}
}
