package connections

import (
	"time"

	"github.com/pion/webrtc/v3"
)

func OnOpen(conn *webrtc.DataChannel) {
	for {
		conn.SendText("sup")
		time.Sleep(time.Second)
	}

}
