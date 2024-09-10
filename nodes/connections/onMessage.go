package connections

import (
	"fmt"

	"github.com/pion/webrtc/v3"
)

func OnMessage(conn *webrtc.DataChannel, msg webrtc.DataChannelMessage) {
	fmt.Println("mensaje recibido", string(msg.Data))

}
