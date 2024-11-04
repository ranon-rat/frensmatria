package connections

import (
	"fmt"
	"time"

	"github.com/fatih/color"
	"github.com/ranon-rat/frensmatria/nodes/src/core"
)

func OnOpen(conn ConnectionID) {
	core.LogColor("New Connection", len(Conns))
	go SendAlive(conn) // with this we only say "hey, i am alive :D"
	if ComparingQ {
		go OnOpenComparing(conn)
	}
	ConnectedNodes++
	CloseIfNoResponse(conn)
}
func OnOpenComparing(conn ConnectionID) {

	ComparingNodes++
	core.LogColor(color.New(color.FgGreen).Sprint("sending event:"), color.New(color.FgHiYellow).Sprint("get"))

	if err := conn.Connection.SendText(fmt.Sprintf("get %d", LastDate)); err != nil {
		conn.Connection.Close()
		return
	}

	LifeTime(10, 30, 5, max(1, 5/len(Conns)), 1, CIncreaseLifeTime[conn.ID])
	OnEnding(conn.ID)
}

func SendAlive(conn ConnectionID) {
	for {
		if err := conn.Connection.SendText("A"); err != nil {
			conn.Connection.Close()
			break
		}
		time.Sleep(time.Second * 5)
	}
}

func CloseIfNoResponse(conn ConnectionID) {
	LifeTime(10, 30, 5, 10, 1, Alive[conn.ID])
	conn.Connection.Close()

	core.LogColor("disconnecting:", color.New(color.Bold, color.FgRed).Sprint("reason timeout"))
	// okay so, for some reason we have some problems when i close the program with ctrl+c
	// so i need to do this shit, fuck
	OnClose(conn)
}

// this works for keeping something alive the connection and in case that it has passed way too much time
// the function just finish, so its useful for things like timeout, and other stuff.

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
