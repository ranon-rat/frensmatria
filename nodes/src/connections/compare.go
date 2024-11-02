package connections

import (
	"fmt"
	"sync"

	"github.com/ranon-rat/frensmatria/nodes/src/core"
	"github.com/ranon-rat/frensmatria/nodes/src/core/channels"
	"github.com/ranon-rat/frensmatria/nodes/src/db"
)

func CompareEndingEvent() {
	for ComparingQ {
		<-CompareEndChan
		if ComparingNodes == 0 {
			CompareResults(ComparingMap, LastDate)
			ComparingMap = map[string]map[string]int{}
			ComparingQ = false
			ComparingQs = map[string]bool{}
		}
	}

}

func CompareResults(compare map[string]map[string]int, currentDate int) {
	final := make(map[string]int)
	// in case each map has different size
	for _, m := range compare {
		for input := range m {
			if _, e := final[input]; e {
				continue
			}

			dates := make(map[int]int)
			newInput := false

			// maybe i will add something for all the nodes to avoid adding things that are not popular
			for _, n := range compare {
				if _, e := n[input]; !e {
					newInput = true
					continue
				}
				d := n[input] // i dont wish for anything old, and in thisi case it should return an error
				if d <= currentDate {
					continue
				}
				dates[d]++
			}
			date := GetBiggerDate(dates)
			if newInput {
				channels.SendMessage(fmt.Sprintf("new %s", core.GematriaSharing2Base64(core.GematriaSharing{
					Content: input,
					Date:    date,
				})), "")
			}
			final[input] = date
		}
	}
	var wg sync.WaitGroup
	for i, d := range final {
		wg.Add(1)
		go func() {
			defer wg.Done()
			db.AddGematria(i, d)
		}()
	}
	wg.Wait()
}

func GetBiggerDate(dates map[int]int) int {
	bigDate := 0
	bigger := 0
	for date, r := range dates {
		if r > bigger {
			bigger = r
			bigDate = date
		}

	}
	return bigDate

}
