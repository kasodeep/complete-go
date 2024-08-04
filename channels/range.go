package channels

/*
The for range loop will get the value from the channel and print it.
It autmatically exists when the channel closes.
*/
func ConcurrentFib(n int) []int {
	chInts := make(chan int)
	series := make([]int, 0, n)

	// run in parallel and get the series.
	go func() {
		fibonacci(n, chInts)
	}()

	for number := range chInts {
		series = append(series, number)
	}
	return series
}

func fibonacci(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
	}
	close(ch)
}
