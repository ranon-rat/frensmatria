package SDPConn

import (
	"github.com/pion/webrtc/v3"
	"github.com/ranon-rat/frensmatria/nodes/connections"
)

/*
this is just the data channel handler,
there is not much to add, since i will
be communicating with the connections package
the code will be the same for incoming connections and
for entering connections
*/

func dcHandler(d *webrtc.DataChannel) {
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
