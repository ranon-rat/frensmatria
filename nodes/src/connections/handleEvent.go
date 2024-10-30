package connections

import (
	"fmt"
	"math/rand"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
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

		// so this will be listening when we close the channel
		go func() {
			OnOpen(conn)
		}()
		go func() {
			for {

				ComparingMap = append(ComparingMap, make(map[string]int))
				ComparingQs = append(ComparingQs, true)
				msg := <-msgChan
				OnMessage(conn, msg, len(ComparingMap)-1)
			}
		}()
		go func() {
			<-closeChan
			OnClose(cID)
			//delete(SDPConn, conn)
		}()

	}
}
