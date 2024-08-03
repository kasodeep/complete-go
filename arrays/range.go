package arrays

import "fmt"

func RangeDemo() {
	fruits := []string{"apple", "banana", "grape"}
	for i, fruit := range fruits {
		fmt.Print(i, " ", fruit, " ")
	}
	fmt.Println()
}
