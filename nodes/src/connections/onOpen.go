package connections

import (
	"fmt"

	"github.com/pion/webrtc/v3"
)

func OnOpen(conn *webrtc.DataChannel) {
	if ComparingQ {
		conn.SendText(fmt.Sprintf("get %d", CurrentDate))
	}
}
