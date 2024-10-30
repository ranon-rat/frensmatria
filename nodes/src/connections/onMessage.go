package connections

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/pion/webrtc/v3"
	"github.com/ranon-rat/frensmatria/nodes/src/compare"
	"github.com/ranon-rat/frensmatria/nodes/src/core"
	"github.com/ranon-rat/frensmatria/nodes/src/core/channels"
	"github.com/ranon-rat/frensmatria/nodes/src/db"
)

// Gematria->connections
// new     base64json // sending or receiving
// {content:"example",date:1234567} // check if the content is already in db
// compare base64json // this is just for sending, or receiving (if you receive this, you shouldnt share it with other nodes)
// {content:"example",date:12341325}
// end (this is for the comparing stuf)
// get dateTime // this is only for getting information
// new, compare, end, get, those are all
func OnMessage(conn *webrtc.DataChannel, msg webrtc.DataChannelMessage, i int) {
	information := strings.Split(string(msg.Data), " ")
	if len(information) < 2 {
		return
	}

	switch information[0] {
	case "get":
		date, _ := strconv.Atoi(information[1])
		db.GetAllGematria(conn, date)

	case "new":
		g := core.Base64_2GematriaSharing(information[1])
		log.Println("New", g.Content)

		// this is actually important :D
		if db.AddGematria(g.Content, g.Date) == nil {
			channels.ConnectionComm <- fmt.Sprintf("new %s", core.GematriaSharing2Base64(g))
		}
	case "compare":
		if !ComparingQs[i] {
			return
		}
		g := core.Base64_2GematriaSharing(information[1])
		log.Println("Comparing", g.Content)
		ComparingMap[i][g.Content] = g.Date

	case "end":

		// not finished yet, i still need to modify some other stuff for improving the system
		if ComparingQs[i] {
			ComparingNodes--
			ComparingQs[i] = false
		}
		if ComparingNodes == 0 {
			compare.Compare(ComparingMap, CurrentDate)
			ComparingMap = []map[string]int{}
			ComparingQ = false
		}

	default:
		return
	}

}

/*
pensemos que vamos a hacer...
*/
