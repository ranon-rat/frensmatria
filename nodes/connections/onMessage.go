package connections

import (
	"fmt"

	"github.com/pion/webrtc/v3"
)

func OnMessage(conn *webrtc.DataChannel, msg webrtc.DataChannelMessage, ID string) {
	fmt.Println("mensaje recibido", string(msg.Data))
	fmt.Printf("\r > %s\n\r>", string(msg.Data))

	SetMSG(string(msg.Data), ID)

}
