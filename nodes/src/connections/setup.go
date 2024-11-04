package connections

import (
	"github.com/ranon-rat/frensmatria/nodes/src/core/channels"
	"github.com/ranon-rat/frensmatria/nodes/src/db"
)

func Setup(update bool) {
	ExpectedNodes = <-channels.HowManyNodes
	if update {
		LastDate = db.GetLastDate()
		ComparingQ = true
		go func() {
			for {
				if ComparingNodes >= ExpectedNodes {
					CompareEndingEvent()
				}
			}
		}()

	}
	go SendMessages()
	go HandleEventConns()

}
