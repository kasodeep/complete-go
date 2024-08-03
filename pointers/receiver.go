package pointers

import "fmt"

type car struct {
	color string
}

func (c *car) setColor(color string) {
	c.color = color
}

func ReceiverDemo() {
	c := car{
		color: "white",
	}
	c.setColor("blue")
	fmt.Println(c.color)
}
