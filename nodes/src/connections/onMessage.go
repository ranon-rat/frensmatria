package connections

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/pion/webrtc/v3"
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
func OnMessage(conn *webrtc.DataChannel, msg webrtc.DataChannelMessage, id string) {
	information := strings.Split(string(msg.Data), " ")
	if len(information) < 2 {
		return
	}
	switch information[0] {
	case "new":
		g := core.Base64_2GematriaSharing(information[1])
		log.Println("New", g.Content)

		// this is actually important :D
		if g.Content == "" {
			return
		}
		if db.AddGematria(g.Content, g.Date) == nil {
			channels.SendMessage(fmt.Sprintf("new %s", core.GematriaSharing2Base64(g)), id)
		}
	case "get":
		date, _ := strconv.Atoi(information[1])
		db.GetAllGematria(conn, date)

	case "compare":
		if !ComparingQs[id] {
			return
		}
		IncreaseLifeTime[id] <- struct{}{}
		g := core.Base64_2GematriaSharing(information[1])
		if g.Content == "" {
			return
		}
		log.Println("Comparing", g.Content)
		ComparingMap[id][g.Content] = g.Date

	case "end":
		// not finished yet, i still need to modify some other stuff for improving the system
		OnEnding(id)

	default:
		return
	}

}

func OnEnding(id string) {
	if ComparingQs[id] {
		ComparingNodes--
		delete(ComparingQs, id)
		CompareEndChan <- struct{}{}
	}
}
