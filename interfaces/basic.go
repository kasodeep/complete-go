// Interfaces are just collections of method signatures.
// A type "implements" an interface if it has methods that match the interface's method signatures.
package interfaces

import "fmt"

type employee interface {
	getName() string
	getSalary() int
}

type contractor struct {
	name         string
	hourlyPay    int
	hoursPerYear int
}

func (c contractor) getName() string {
	return c.name
}

func (c contractor) getSalary() int {
	return c.hourlyPay * c.hoursPerYear
}

type fullTime struct {
	name   string
	salary int
}

func (ft fullTime) getSalary() int {
	return ft.salary
}

func (ft fullTime) getName() string {
	return ft.name
}

func InterfaceDemo() {
	c1 := contractor{name: "John Doe", hourlyPay: 50, hoursPerYear: 2000}
	ft1 := fullTime{name: "Jane Smith", salary: 90000}

	employees := []employee{c1, ft1}

	// iterate through the slice and print the name and salary of each employee.
	for _, emp := range employees {
		fmt.Printf("Name: %s, Salary: $%d\n", emp.getName(), emp.getSalary())
	}
	fmt.Println()
}
