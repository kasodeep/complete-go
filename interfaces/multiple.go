package interfaces

import "fmt"

func (e email) cost() float64 {
	if !e.isSubscribed {
		return 0.05 * float64(len(e.body))
	}
	return 0.01 * float64(len(e.body))
}

func (e email) format() string {
	return fmt.Sprintf(e.body)
}

type expense interface {
	cost() float64
}

type formatter interface {
	format() string
}

type email struct {
	isSubscribed bool
	body         string
}

func MultipleDemo(e email) {
	var exp expense = e
	var fmttr formatter = e

	fmt.Printf("Email Cost: $%.2f\n", exp.cost())
	fmt.Printf("Email Body: %s\n", fmttr.format())
}
