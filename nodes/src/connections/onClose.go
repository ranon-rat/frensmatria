package connections

import "log"

func OnClose(conn ConnectionID) {
	delete(Conns, conn)
	log.Println("Closing Connection")
	if ComparingQ {
		delete(ComparingMap, conn.ID)
		delete(ComparingQs, conn.ID)
		delete(IncreaseLifeTime, conn.ID)

	}
}
