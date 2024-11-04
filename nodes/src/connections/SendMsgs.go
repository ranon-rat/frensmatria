package connections

import (

	"github.com/ranon-rat/frensmatria/nodes/src/core/channels"
)

func SendMessages() {

	for {
		content := <-channels.ConnectionComm
		for v := range Conns {
			if v.ID == content.ID {
				continue
			}
			if err := v.Connection.SendText(content.Content); err != nil {
				v.Connection.Close()

				continue
			}
		}
	}
}
