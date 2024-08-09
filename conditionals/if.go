package conditionals

import "fmt"

func getLength(s string) int {
	return len(s)
}

func IfDemo() {
	messageLen := 10
	maxMessageLen := 20
	fmt.Println("Trying to send a message of length:", messageLen, "and a max length of:", maxMessageLen)

	if messageLen <= maxMessageLen {
		fmt.Println("Message sent")
	} else {
		fmt.Println("Message not sent")
	}

	email := "This is an email"

	// if scoped variable.
	if length := getLength(email); length < 1 {
		fmt.Println("Email is invalid")
	}
	fmt.Println()
}
