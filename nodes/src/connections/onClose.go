package connections

import "github.com/ranon-rat/frensmatria/nodes/src/core"

func OnClose(conn ConnectionID) {
	core.LogColor("Closing Connection")
	delete(Conns, conn)
	if ComparingQ {
		delete(ComparingMap, conn.ID)
		delete(ComparingQs, conn.ID)
		delete(IncreaseLifeTime, conn.ID)
	}
}
