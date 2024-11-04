package connections

import (
	"fmt"

	"github.com/ranon-rat/frensmatria/nodes/src/core/channels"
	"github.com/ranon-rat/frensmatria/nodes/src/db"
)

func Setup(update bool) {
	ExpectedNodes = <-channels.HowManyNodes
	fmt.Println(ExpectedNodes, ConnectedNodes)
	if update {
		LastDate = db.GetLastDate()
		ComparingQ = true
		go func() {
			for {
				if ConnectedNodes >= ExpectedNodes {
					fmt.Println(ConnectedNodes)
					CompareEndingEvent()
					return
				}
			}
		}()

	}
	go SendMessages()
	go HandleEventConns()

}
