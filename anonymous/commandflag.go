package anonymous

import (
	"flag"
	"fmt"
)

// to run this `go run commandflag.go -word=bar -numb=7 -fork -svar=flag`
func FlagsDemo() {
	// it creates a pointer to flag with name, default value and description.
	wordPtr := flag.String("word", "foo", "a string")
	numbPtr := flag.Int("numb", 42, "an int")
	forkPtr := flag.Bool("fork", false, "a bool")

	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")
	flag.Parse()

	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *forkPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())
}
