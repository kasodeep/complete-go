package channels

import (
	"time"
)

/*
Channels are a typed, thread-safe queue.
Channels allow different goroutines to communicate with each other.
*/

/*
The channel is a blocking operation if the channel is empty or full.
*/

/*
We have read only and write only channels.
The type is specified in the function signature.
(name <-chan int) reader
(name chan<- int) writer
*/

type email struct {
	date time.Time
}

func CheckEmailAge(emails [3]email) [3]bool {
	isOldChan := make(chan bool)

	// without go there will be a deadlocked cause a sender needs a receiver on the other end.
	go sendIsOld(isOldChan, emails)

	isOld := [3]bool{}
	isOld[0] = <-isOldChan
	isOld[1] = <-isOldChan
	isOld[2] = <-isOldChan
	return isOld
}

// don't touch below this line

func sendIsOld(isOldChan chan<- bool, emails [3]email) {
	for _, e := range emails {
		if e.date.Before(time.Date(2020, 0, 0, 0, 0, 0, 0, time.UTC)) {
			isOldChan <- true
			continue
		}
		isOldChan <- false
	}
}
