package errors

import (
	"errors"
	"fmt"
)

func personalDivide(x, y float64) (float64, error) {
	if y == 0 {
		return 0, errors.New("no dividing by 0")
	}
	return x / y, nil
}

func ErrorPackage() {
	_, err := personalDivide(10, 0)
	if err != nil {
		fmt.Println(err)
	}
}
