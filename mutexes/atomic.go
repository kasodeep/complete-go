package mutexes

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func AtomicDemo() {

	// atomic variable coming from the sync package.
	var ops atomic.Uint64
	var wg sync.WaitGroup

	// creating 50 worker threads.
	for i := 0; i < 50; i++ {
		wg.Add(1)

		// running the thread on separate go routine.
		go func() {
			for c := 0; c < 1000; c++ {
				ops.Add(1)
			}
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("ops:", ops.Load())
}
