package channels

import "fmt"

func Index() {
	fmt.Println("----- Channels Demo -----")

	SendEmail("Hello World!")
	fmt.Println(ConcurrentFib(10))

	fmt.Println()
	TimeoutDemo()
}
