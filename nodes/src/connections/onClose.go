package connections

func OnClose(conn ConnectionID) {
	delete(Conns, conn)
}
