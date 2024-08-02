package main

import "fmt"

func main() {
	// Constants are declared with the const keyword. 
	// They can't use the := short declaration syntax.
	const premiumPlanName = "Premium Plan"
	const basicPlanName = "Basic Plan"

	fmt.Println("plan:", premiumPlanName)
	fmt.Println("plan:", basicPlanName)
}
