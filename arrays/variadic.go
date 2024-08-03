package arrays

import "fmt"

// A variadic function receives the variadic arguments as a slice.
func concat(strs ...string) string {
	final := ""
	for i := 0; i < len(strs); i++ {
		final += strs[i]
	}
	return final
}

func printStrings(strings ...string) {
	for i := 0; i < len(strings); i++ {
		fmt.Print(strings[i], " ")
	}
	fmt.Println()
}

func VariadicDemo() {
	fmt.Println("Variadic Demo:")
	final := concat("Hello ", "there ", "friend!")
	fmt.Println(final)

	// spread operator.
	names := []string{"bob", "sue", "alice"}
	printStrings(names...)
}
