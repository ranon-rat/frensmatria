package SDPConn

import (
	"github.com/pion/webrtc/v3"
	"github.com/ranon-rat/frensmatria/nodes/src/connections"
)

/*
this is just the data channel handler,
there is not much to add, since i will
be communicating with the connections package
the code will be the same for incoming connections and
for entering connections
*/

func dcHandler(d *webrtc.DataChannel) {

	d.OnOpen(func() {
		connections.ConnChan <- d

	})

}
