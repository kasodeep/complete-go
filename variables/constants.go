package variables

import "fmt"

func Constants() {
	// constants are declared with the const keyword.
	// they can't use the := short declaration syntax.
	const premiumPlanName = "Premium Plan"
	const basicPlanName = "Basic Plan"

	fmt.Println("Plan:", premiumPlanName)
	fmt.Println("Plan:", basicPlanName)
	fmt.Println()
}
