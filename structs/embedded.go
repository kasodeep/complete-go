package structs

import "fmt"

/*
	"car" is embedded, so the definition of a "truck" now also additionally, contains all of the fields of the car struct.
*/
type truck struct {
	car
	bedSize int
}

func EmbeddedDemo() {
	lanesTruck := truck{
		bedSize: 10,
		car: car{
			make:  "toyota",
			model: "camry",
		},
	}

	// embedded fields promoted to the top-level instead of lanesTruck.car.make
	fmt.Printf("Embedded Fields (Make & BedSize): %s %d \n", lanesTruck.make, lanesTruck.bedSize)
}
