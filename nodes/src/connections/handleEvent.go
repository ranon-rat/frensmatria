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
		ID := RandStringRunes(10)

		fmt.Println("sup we are back")
		cID := ConnectionID{
			ID:         ID,
			Connection: conn,
		}
		Conns[cID] = true
		if ComparingQ {

			IncreaseLifeTime[ID] = make(chan struct{})
		}
		// so this will be listening when we close the channel
		go func() {
			OnOpen(conn, ID)
		}()
		go func() {
			for {
				if ComparingQ {
					ComparingMap[ID] = make(map[string]int)
					ComparingQs[ID] = true
				}

				msg := <-msgChan
				OnMessage(conn, msg, ID)
			}
		}()
		go func() {
			<-closeChan
			OnClose(cID)
		}()

	}
}
