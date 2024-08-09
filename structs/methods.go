package structs

import "fmt"

type rect struct {
	width  int
	height int
}

// The method area has a receiver of (r rect) rect is the struct r is the placeholder.
func (r *rect) area() int {
	return r.width * r.height
}

func StructMethods() {
	var r = rect{
		width:  5,
		height: 10,
	}

	fmt.Println("Sturct Method:", r.area())
	fmt.Println()
}
