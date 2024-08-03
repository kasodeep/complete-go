package pointers

type Analytics struct {
	MessagesTotal     int
	MessagesFailed    int
	MessagesSucceeded int
}

type Message struct {
	Recipient string
	Success   bool
}

func GetMessageText(a *Analytics, m Message) {
	a.MessagesTotal++
	if m.Success {
		a.MessagesSucceeded++
	} else {
		a.MessagesFailed++
	}
}
