package connections

import (
	"fmt"

	"github.com/ranon-rat/frensmatria/core"
)

func HandleEventConns() {

	for {

		connInfo := <-ConnInfoChan
		closeChan := connInfo.CloseChan
		msgChan := connInfo.MsgChan
		conn := connInfo.Connection
		ID := core.RandStringRunes(10)

		fmt.Println("sup we are back")
		cID := ConnectionID{
			ID:         ID,
			Connection: conn,
		}
		Conns[cID] = true

		// so this will be listening when we close the channel
		go func() {
			OnOpen(conn)
		}()
		go func() {
			for {

				msg := <-msgChan
				OnMessage(conn, msg, ID)
			}
		}()
		go func() {
			<-closeChan
			OnClose(cID)
			//delete(SDPConn, conn)
		}()

	}
}
