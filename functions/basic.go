package functions

import "fmt"

// We need to use the keyword func, along with type safety.
// The type comes after the variable.
func concat(s1 string, s2 string) string {
	return s1 + s2
}

// hp & damage have the same type int.
// When the same type parameter are next to each other.
func AddToDatabase(hp, damage int, name string, level int) {
	newHp := damage * level
	fmt.Println("Name:", name, "New HP:", newHp)
}

func test(s1 string, s2 string) {
	fmt.Println(concat(s1, s2))
	fmt.Println()
}

func FunctionsDemo() {
	test("Lane,", " Happy birthday!")
}
