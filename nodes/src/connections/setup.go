package connections

func Setup() {
	SetDate()
	go CompareEndingEvent()
	go SendMessages()
	HandleEventConns()
}
