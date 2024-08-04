package mutexes

import (
	"fmt"
	"sync"
	"time"
)

/*
Mutexes allow us to lock access to data.
This ensures that we can control which goroutines can access certain data at which time.
*/
type safeCounter struct {
	counts map[string]int
	mux    *sync.Mutex
}

func (sc safeCounter) inc(key string) {
	sc.mux.Lock()
	defer sc.mux.Unlock()
	sc.slowIncrement(key)
}

func (sc safeCounter) val(key string) int {
	sc.mux.Lock()
	defer sc.mux.Unlock()
	return sc.slowVal(key)
}

func (sc safeCounter) slowIncrement(key string) {
	tempCounter := sc.counts[key]
	time.Sleep(time.Microsecond)
	tempCounter++
	sc.counts[key] = tempCounter
}

func (sc safeCounter) slowVal(key string) int {
	time.Sleep(time.Microsecond)
	return sc.counts[key]
}

func MutexDemo() {
	counter := safeCounter{counts: make(map[string]int), mux: &sync.Mutex{}}
	for i := 0; i < 10; i++ {
		go counter.inc("somekey")
	}

	time.Sleep(time.Millisecond)
	fmt.Println(counter.val("somekey"))
}
