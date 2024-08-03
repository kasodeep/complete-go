package variables

import "fmt"

func ShortNotation() {
	// short declaration, the := operator is used for short variable declaration.
	messageStart := "Happy birthday! You are now"
	age := 21
	messageEnd := "years old!"
	fmt.Println(messageStart, age, messageEnd)

	// same line declaration.
	averageOpenRate, displayMessage := .23, "is the average open rate of your messages"
	fmt.Println(averageOpenRate, displayMessage)
}
