package connections

func Setup(update bool) {
	if update {
		SetDate()
	}
	go CompareEndingEvent()
	go SendMessages()
	HandleEventConns()
}
