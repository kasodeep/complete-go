package channels

import (
	"fmt"
	"time"
)

/*
Go was designed to be concurrent, which is a trait fairly unique to Go.
It excels at performing many tasks simultaneously safely using a simple syntax.
*/
func SendEmail(message string) {
	// this function will be executed in a goroutine. we can't expect a return value.
	go func() {
		time.Sleep(time.Millisecond * 250)
		fmt.Printf("Email received: '%s'\n", message)
	}()

	fmt.Printf("Email sent: '%s'\n", message)
	time.Sleep(500 * time.Millisecond)
}
