package connections

import "github.com/ranon-rat/frensmatria/nodes/src/core/channels"

func SendMessages() {

	for {
		// this will send anything related to anything that i am interested in
		content := <-channels.ConnectionComm
		for v := range Conns {

			if err := v.Connection.SendText(content); err != nil {
				delete(Conns, v)
				continue
			}

		}
	}
}
