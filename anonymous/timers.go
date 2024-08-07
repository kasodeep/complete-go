package anonymous

import (
	"fmt"
	"time"
)

func TimersDemo() {
	// this creates a channel that will fire after 2 seconds.
	timer1 := time.NewTimer(2 * time.Second)

	// this blocks until the timer fires.
	<-timer1.C
	fmt.Println("Timer 1 fired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()

	// not allowing timer to fire.
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

	time.Sleep(2 * time.Second)
}
