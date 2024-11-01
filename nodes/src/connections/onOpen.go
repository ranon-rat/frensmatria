package connections

import (
	"fmt"
	"time"

	"github.com/pion/webrtc/v3"
)

func OnOpen(conn *webrtc.DataChannel, ID string) {
	if ComparingQ {
		ComparingNodes++

		conn.SendText(fmt.Sprintf("get %d", LastDate))
	}
	if !ComparingQ {
		return
	}
	lifeTime := 10
	go func() {
		for {
			<-IncreaseLifeTime[ID]
			lifeTime += 5
		}
	}()
	for lifeTime > 0 {
		time.Sleep(time.Second * 2)
		lifeTime--
	}
	if ComparingQs[ID] {
		ComparingNodes--
		delete(ComparingQs, ID)
		CompareEndChan <- struct{}{}

	}
}
