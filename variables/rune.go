package variables

import (
	"fmt"
	"unicode/utf8"
)

/*
Go also has a special type, rune, which is an alias for int32.
This means that a rune is a 32-bit integer, which is large enough to hold any Unicode code point.
*/
func RuneDemo() {
	const name = "🐻"
	fmt.Printf("'name' byte length: %d\n", len(name))                    // 4
	fmt.Printf("'name' rune length: %d\n", utf8.RuneCountInString(name)) // 1
}
