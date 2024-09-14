package connections

func SendMessages() {

	for {
		msg := <-MsgChan

		for v := range Conns {
			if msg.ID == v.ID {
				continue
			}
			if err := v.Connection.SendText(msg.Content); err != nil {
				delete(Conns, v)
				continue
			}

		}
	}
}
