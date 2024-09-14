package SDPConn

func Setup() {
	go IncomingConn()
	go EnterConn()
	go ConnectToNodes()
}
