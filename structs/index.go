package structs

import "fmt"

func Index() {
	fmt.Println("----- Structs Demo -----")

	// basic struct demo.
	StructsDemo()
	Anonymous()

	// embedded and methods demo.
	EmbeddedDemo()
	StructMethods()

	// size and empty demo.
	SizeDemo()
	EmptyStruct()
}
