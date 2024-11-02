package connections

import (
	"fmt"
	"log"
	"time"

	"github.com/pion/webrtc/v3"
)

func OnOpen(conn *webrtc.DataChannel, ID string) {
	log.Println("New Connection")
	if !ComparingQ {
		return
	}
	ComparingNodes++
	conn.SendText(fmt.Sprintf("get %d", LastDate))

	lifeTime := 10
	go func() {
		for {
			<-IncreaseLifeTime[ID]
			// i will just wait a little :D
			lifeTime += max(1, 10/len(Conns))
			lifeTime = min(lifeTime, 50)
		}
	}()
	for lifeTime > 0 {
		time.Sleep(time.Second * 2)
		lifeTime--
	}
	OnEnding(ID)

}
