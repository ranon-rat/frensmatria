package connections

func SetMSG(msg, id string) {
	MsgChan <- Message{Content: msg, ID: id}
}
