package SDPConn

func Setup() {
	go OfferSDPConn()
	go ONAnswer()
	go ConnectToNodes()
}
