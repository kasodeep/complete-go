package interfaces

import "fmt"

func printNumericValue(num interface{}) {
	switch v := num.(type) {
	case int:
		fmt.Printf("%T\n", v)
	case string:
		fmt.Printf("%T\n", v)
	default:
		fmt.Printf("%T\n", v)
	}
}

func StructInterfaceTypes() {
	printNumericValue(1)
	printNumericValue("1")
	printNumericValue(struct{}{})
}
