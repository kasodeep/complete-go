package conditionals

import "fmt"

/*
	The switch statement in Go doesn't have break, it is default setting.
	We can use `fallthrough` if we still want to go to next case.
*/
func billingCost(plan string) float64 {
	switch plan {
	case "basic":
		return 10.0
	case "pro":
		return 20.0
	case "enterprise":
		return 50.0
	default:
		return 0.0
	}
}

func SwitchDemo() {
	plan := "pro"
	fmt.Printf("The cost for a %s plan is $%.2f\n", plan, billingCost(plan))
}
