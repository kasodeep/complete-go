package anonymous

import (
	"fmt"
	"sync"
	"time"
)

func work(id int) {
	fmt.Printf("Worker %d starting\n", id)

	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func WaitGroupDemo() {

	// this is like countdown latch in java.
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done() // make sure to decrement the count.
			work(i)
		}()
	}

	wg.Wait()
}
