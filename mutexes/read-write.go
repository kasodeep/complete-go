package mutexes

import (
	"fmt"
	"sync"
	"time"
)

type safeCounterTwo struct {
	counts map[string]int
	mux    *sync.RWMutex
}

func (sc safeCounterTwo) inc(key string) {
	sc.mux.Lock()
	defer sc.mux.Unlock()
	sc.slowIncrement(key)
}

func (sc safeCounterTwo) val(key string) int {
	sc.mux.RLock()
	defer sc.mux.RUnlock()
	return sc.counts[key]
}

func (sc safeCounterTwo) slowIncrement(key string) {
	tempCounter := sc.counts[key]
	time.Sleep(time.Microsecond)
	tempCounter++
	sc.counts[key] = tempCounter
}

func ReadWriteDemo() {
	counter := safeCounterTwo{counts: make(map[string]int), mux: &sync.RWMutex{}}
	for i := 0; i < 10; i++ {
		go counter.inc("somekey")
	}

	time.Sleep(time.Millisecond)
	fmt.Println("Count of somekey:", counter.val("somekey"))
}
