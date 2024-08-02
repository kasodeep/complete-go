package main

import "fmt"

func printReports(intro, body, outro string) {
	// Function with no name and defining body while passing.
	printCostReport(func(a string) int {
		return 2 * len(intro)
	}, intro)

	printCostReport(func(a string) int {
		return 3 * len(body)
	}, body)
	
	printCostReport(func(a string) int {
		return 4 * len(outro)
	}, outro)
}

func main() {
	printReports(
		"Welcome to the Hotel California",
		"Such a lovely place",
		"Plenty of room at the Hotel California",
	)
}

func printCostReport(costCalculator func(string) int, message string) {
	cost := costCalculator(message)
	fmt.Printf(`Message: "%s" Cost: %v cents`, message, cost)
	fmt.Println()
}
