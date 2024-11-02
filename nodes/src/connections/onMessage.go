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
// new     base64json // sending or receiving, it will check that if its in the db
// compare base64json // this is just for sending, or receiving (if you receive this, you shouldnt share it with other nodes)
// end (this is for the comparing stuf)
// get dateTime // this is only for getting information
// new, compare, end, get, those are all

// probably i will add something new for the messages
func OnMessage(conn ConnectionID, msg webrtc.DataChannelMessage) {
	ID := conn.ID
	information := strings.Split(string(msg.Data), " ")
	if len(information) < 2 {
		return
	}
	switch information[0] {
	case "new":
		g := core.Base64_2GematriaSharing(information[1])
		if g.Content == "" {
			return
		}
		log.Println(information[0], g.Content, g.Date)
		if db.AddGematria(g.Content, g.Date) == nil {
			channels.SendMessage(fmt.Sprintf("new %s", core.GematriaSharing2Base64(g)), ID)
		}
	case "get":
		date, _ := strconv.Atoi(information[1])
		if date == 0 {
			return
		}
		log.Println(information[0], date)
		if db.GetAllGematria(conn.Connection, date) != nil {
			OnClose(conn)
		}
	case "compare":
		if !ComparingQs[ID] {
			// maybe i should add something so it just stops?
			return
		}
		IncreaseLifeTime[ID] <- struct{}{}
		g := core.Base64_2GematriaSharing(information[1])
		if g.Content == "" {
			return
		}
		log.Println(information[0], g.Content, g.Date)
		ComparingMap[ID][g.Content] = g.Date

	case "end":
		log.Println(information[0])
		OnEnding(ID)
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
