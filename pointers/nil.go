package pointers

import (
	"strings"
)

/*
If a pointer points to nothing (the zero value of the pointer type),
then dereferencing it will cause a runtime error (a panic) that crashes the program.
*/
func RemoveProfanity(message *string) {
	if message == nil {
		return
	}

	messageVal := *message
	messageVal = strings.ReplaceAll(messageVal, "fubb", "****")
	messageVal = strings.ReplaceAll(messageVal, "shiz", "****")
	messageVal = strings.ReplaceAll(messageVal, "witch", "*****")
	*message = messageVal
}
