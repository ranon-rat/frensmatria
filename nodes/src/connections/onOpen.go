package connections

import (
	"fmt"
	"time"

	"github.com/ranon-rat/frensmatria/nodes/src/core"
)

func OnOpen(conn ConnectionID) {
	core.LogColor("New Connection")
	if !ComparingQ {
		go OnOpenComparing(conn)
	}
	go SendAlive(conn)
	CloseIfNoResponse(conn)
}
func OnOpenComparing(conn ConnectionID) {
	ComparingNodes++
	conn.Connection.SendText(fmt.Sprintf("get %d", LastDate))

	LifeTime(10, 30, 5, max(1, 10/len(Conns)), 1, IncreaseLifeTime[conn.ID])
	OnEnding(conn.ID)
}
func SendAlive(conn ConnectionID) {
	for {
		if err := conn.Connection.SendText("A"); err != nil {
			conn.Connection.Close()
			break
		}
		time.Sleep(time.Second * 15)
	}
}
func CloseIfNoResponse(conn ConnectionID) {
	LifeTime(10, 30, 5, 4, 1, Alive[conn.ID])
	conn.Connection.Close()
	core.LogColor("disconnecting because of no response")
	OnClose(conn)
}

// the initial value which it starts with
// the maximum value
// interval
// addition of the liftime and the subtraction of the liftime
func LifeTime(initialLifeTime, maxVal, interval, adding, substract int, check chan struct{}) {
	lifeTime := initialLifeTime
	go func() {
		for {
			<-check
			// i will just wait a little :D
			lifeTime += adding
			lifeTime = min(lifeTime, maxVal)
		}
	}()
	for lifeTime > 0 {
		time.Sleep(time.Second * time.Duration(interval))
		lifeTime -= substract
	}
}
