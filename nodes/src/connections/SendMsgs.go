package connections

import "github.com/ranon-rat/frensmatria/nodes/src/core/channels"

func SendMessages() {

	for {
		// this will send anything related to anything that i am interested in
		content := <-channels.ConnectionComm
		SendMessageEveryone(content)
	}
}
func SendMessageEveryone(content channels.Message) {
	for v := range Conns {
		if v.ID == content.ID {
			continue
		}
		if err := v.Connection.SendText(content.Content); err != nil {
			delete(Conns, v)
			continue
		}
	}
}
