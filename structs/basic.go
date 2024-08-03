package structs

import "fmt"

// struct car.
type car struct {
	make    string
	model   string
	doors   int
	mileage int

	// nested structs.
	frontWheel wheel
	backWheel  wheel
}

// struct wheel.
type wheel struct {
	radius   int
	material string
}

func StructsDemo() {
	myCar := car{}
	myCar.frontWheel.radius = 5

	newCar := car{
		make:    "Ford",
		model:   "Mustang",
		doors:   2,
		mileage: 20,
		frontWheel: wheel{
			radius:   15,
			material: "rubber",
		},
		backWheel: wheel{
			radius:   10,
			material: "aluminum",
		},
	}

	fmt.Println(newCar)
}
