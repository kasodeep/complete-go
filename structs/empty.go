package structs

import "fmt"

func EmptyStruct() {
	// anonymous empty struct type
	empty := struct{}{}

	// named empty struct type
	type emptyStruct struct{}
	empty = emptyStruct{}

	fmt.Println(empty)
}
