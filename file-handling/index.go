package handling

import "fmt"

func Index() {
	fmt.Println("----- File Handling Demo -----")

	fmt.Println("Paths Demo:")
	PathsDemo()
	fmt.Println()

	fmt.Println("Writing Demo:")
	WritingDemo()
	fmt.Println()

	fmt.Println("Reading Demo:")
	ReadingDemo()
}
