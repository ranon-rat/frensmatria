package connections

import (
	"fmt"
	"time"

	"github.com/pion/webrtc/v3"
)

func OnOpen(conn *webrtc.DataChannel, ID string) {
	if ComparingQ {
		conn.SendText(fmt.Sprintf("get %d", LastDate))
	}

	lifeTime := 10
	go func() {
		for {
			<-IncreaseLifeTime[ID]
			lifeTime += 5
		}
	}()
	for lifeTime > 0 {

		time.Sleep(time.Second)
		lifeTime--
	}
	if ComparingQs[ID] {
		ComparingNodes--
		delete(ComparingQs, ID)
		CompareEndChan <- struct{}{}

	}
}
