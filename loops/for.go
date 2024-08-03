package loops

import "fmt"

/*
Loops in Go can omit sections of a for loop.
For example, the CONDITION (middle part) can be omitted which causes the loop to run forever.
Go does support the continue and break statements.
*/
func ForDemo() {
	for i := 1; i <= 10; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()

	// Go has no while loop.
	num := 0
	sum := 0
	for num <= 10 {
		sum += num
		num++
	}
	fmt.Println("Sum of 10 natural numbers:", sum)
}
