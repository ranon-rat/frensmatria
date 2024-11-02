package connections

import "github.com/ranon-rat/frensmatria/nodes/src/db"

func Setup(update bool) {
	if update {
		LastDate = db.GetLastDate()
		ComparingQ = true
		go CompareEndingEvent()

	}
	go SendMessages()
	go HandleEventConns()

}
