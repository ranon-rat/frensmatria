package connections

import (
	"fmt"
)

func HandleEventConns() {

	for {

		connInfo := <-ConnInfoChan
		closeChan := connInfo.CloseChan
		msgChan := connInfo.MsgChan
		conn := connInfo.Connection

		fmt.Println("sup we are back")
		//SDPConn[conn] = true

		// so this will be listening when we close the channel
		go func() {
			OnOpen(conn)
		}()
		go func() {
			for {

				msg := <-msgChan
				OnMessage(conn, msg)
			}
		}()
		go func() {
			<-closeChan
			OnClose(conn)
			//delete(SDPConn, conn)
		}()

	}
}
