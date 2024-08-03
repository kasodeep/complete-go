package pointers

import "fmt"

/*
A pointer is a variable that stores the memory address of another variable.
This means that a pointer "points to" the location of where the data is stored NOT the actual data itself.
*/

func PointersDemo() {
	// declaring pointer and its variable.
	var p *int
	fmt.Println(p)

	myString := "hello"
	myStringPtr := &myString
	fmt.Println(myStringPtr)

	fmt.Println(*myStringPtr)
	*myStringPtr = "world"
	fmt.Println(myString)
}
