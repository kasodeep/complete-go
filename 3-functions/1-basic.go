package main

import "fmt"

// We need to use the keyword func, along with type safety.
// The type comes after the variable.
func concat(s1 string, s2 string) string {
	return s1 + s2
}

// hp & damage have the same type int.
// When the same type parameter are next to each other.
func addToDatabase(hp, damage int, name string, level int) {
	// ?
}

func test(s1 string, s2 string) {
	fmt.Println(concat(s1, s2))
}

func main() {
	test("Lane,", " happy birthday!")
	test("Elon,", " hope that Tesla thing works out")
	test("Go", " is fantastic")
}
