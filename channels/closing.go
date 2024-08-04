package channels

func CountReports(numSentCh chan int) int {
	total := 0
	for {
		// checking if the channel is closed.
		numSent, ok := <-numSentCh
		if !ok {
			break
		}
		total += numSent
	}
	return total
}

func SendReports(numBatches int, ch chan int) {
	for i := 0; i < numBatches; i++ {
		numReports := i*23 + 32%17
		ch <- numReports
	}
	// the channel should be closed by sender and a buffered channel will be closed when it becomes empty.
	close(ch)
}
