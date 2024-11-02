package connections

import (
	"github.com/fatih/color"
	"github.com/ranon-rat/frensmatria/nodes/src/core"
)

func OnClose(conn ConnectionID) {
	if !Conns[conn] {
		return
	}
	core.LogColor(color.New(color.Bold, color.FgRed).Sprint("Closing Connection"))
	delete(Conns, conn)
	if ComparingQ {
		delete(ComparingMap, conn.ID)
		delete(ComparingQs, conn.ID)
		delete(IncreaseLifeTime, conn.ID)
	}
}
