package errors

import "fmt"

type divideError struct {
	dividend float64
}

// Implementing the error interface.
func (d divideError) Error() string {
	return fmt.Sprintf("can not divide %.1f by zero", d.dividend)
}

// Divide function.
func divide(dividend, divisor float64) (float64, error) {
	if divisor == 0 {
		return 0, divideError{dividend: dividend}
	}
	return dividend / divisor, nil
}

func CustomError() {
	_, err := divide(10, 0)
	if err != nil {
		fmt.Println(err)
	}
}
