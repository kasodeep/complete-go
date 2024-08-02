package main

import "fmt"

func main() {
	fname := "Dalinar"
	lname := "Kholin"
	age := 45
	
	messageRate := 0.5
	isSubscribed := false
	message := "Sometimes a hypocrite is nothing more than a man in the process of changing."
	
	// Sprintf returns the formatted string while Print prints it.
	userLog := fmt.Sprintf("Name: %s %s, Age: %d, Rate: %.1f, isSubscribed: %t, Message: %s", fname, lname, age, messageRate, isSubscribed, message)
	fmt.Println(userLog)
}
