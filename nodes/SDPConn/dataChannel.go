package SDPConn

import (
	"github.com/pion/webrtc/v3"
	"github.com/ranon-rat/frensmatria/nodes/connections"
)

func DataChannelHandler(d *webrtc.DataChannel) {
	closeChan := make(chan struct{})
	msgChan := make(chan webrtc.DataChannelMessage)
	d.OnOpen(func() {
		connections.ConnInfoChan <- connections.ConnectionInfo{
			CloseChan:  closeChan,
			MsgChan:    msgChan,
			Connection: d,
		}

	})
	d.OnClose(func() {
		closeChan <- struct{}{}

	})
	d.OnMessage(func(msg webrtc.DataChannelMessage) {
		msgChan <- msg
	})
}
