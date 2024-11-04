package connections

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/pion/webrtc/v3"
	"github.com/ranon-rat/frensmatria/nodes/src/controllers"
	"github.com/ranon-rat/frensmatria/nodes/src/core"
	"github.com/ranon-rat/frensmatria/nodes/src/core/channels"
	"github.com/ranon-rat/frensmatria/nodes/src/db"
)

// Gematria->connections
// new     base64json // sending or receiving, it will check that if its in the db
// compare base64json // this is just for sending, or receiving (if you receive this, you shouldnt share it with other nodes)
// end (this is for the comparing stuf)
// get dateTime // this is only for getting information
// message
// ... // its for checking that the connection its still up
// new, compare, end, get,message those are all

// probably i will add something new for the messages

func OnMessage(conn ConnectionID, msg webrtc.DataChannelMessage) {

	ID := conn.ID
	Alive[ID] <- struct{}{}
	information := strings.Split(string(msg.Data), " ")
	if len(information) < 2 {
		if information[0] == "end" {
			core.LogColor(color.New(color.FgGreen).Sprint("event:"), color.New(color.FgHiYellow).Sprint(information[0]))
			OnEnding(ID)
		}
		return
	}

	switch information[0] {
	case "new":
		g := core.Base64_2Object[core.GematriaSharing](information[1])
		if g.Content == "" {
			return
		}
		core.LogColor(color.New(color.FgGreen).Sprint("event:"), color.New(color.FgHiYellow).Sprint(information[0]), g.Content)
		if db.AddGematria(g.Content, g.Date) == nil {
			core.LogColor(color.New(color.FgGreen).Sprint("sending event:"), color.New(color.FgHiYellow).Sprint("new"), g.Content)

			channels.SendMessage(fmt.Sprintf("new %s", core.Object2Base64(g)), ID)
		}
	case "get":
		date, _ := strconv.Atoi(information[1])
		if date == 0 {
			return
		}
		core.LogColor(color.New(color.FgGreen).Sprint("event:"), color.New(color.FgHiYellow).Sprint(information[0]), time.Unix(int64(date), 0).Format(time.RFC3339))
		if db.GetAllGematria(conn.Connection, date) != nil {
			conn.Connection.Close()
		}
	case "compare":
		if !ComparingQs[ID] {
			return
		}
		CIncreaseLifeTime[ID] <- struct{}{}
		g := core.Base64_2Object[core.GematriaSharing](information[1])
		if g.Content == "" {
			return
		}
		core.LogColor(color.New(color.FgGreen).Sprint("event:"), color.New(color.FgHiYellow).Sprint(information[0]), g.Content)
		ComparingMap[ID][g.Content] = g.Date
	case "message":
		if MsgCache[information[1]] {
			return
		}
		msg := core.Base64_2Object[core.Messages](information[1])
		controllers.Message <- msg
		channels.SendMessage(fmt.Sprintf("message %s", information[1]), ID)
		MsgCache[information[1]] = true
		time.Sleep(time.Minute * 5)
		delete(MsgCache, information[1])

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
