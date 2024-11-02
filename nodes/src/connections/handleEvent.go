package connections

func HandleEventConns() {

	for {

		connInfo := <-ConnInfoChan
		closeChan := connInfo.CloseChan
		msgChan := connInfo.MsgChan
		conn := connInfo.Connection
		ID := RandStringRunes(10)
		cID := ConnectionID{
			ID:         ID,
			Connection: conn,
		}
		Conns[cID] = true
		if ComparingQ {
			IncreaseLifeTime[ID] = make(chan struct{})
			ComparingMap[ID] = make(map[string]int)
			ComparingQs[ID] = true
		}
		// so this will be listening when we close the channel
		go func() {
			OnOpen(conn, ID)
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
		}()

	}
}
