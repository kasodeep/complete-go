package anonymous

import (
	"bytes"
	"fmt"
	"regexp"
)

func RegexDemo() {
	// match string matches the pattern.
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)

	// otherwise we need to compile the pattern first.
	r, _ := regexp.Compile("p([a-z]+)ch")

	// finding the match, index and string.
	fmt.Println(r.MatchString("peach"))
	fmt.Println(r.FindString("peach punch"))
	fmt.Println("idx:", r.FindStringIndex("peach punch"))

	// sub match divides the pattern into sub groups and compares.
	fmt.Println(r.FindStringSubmatch("peach punch"))
	fmt.Println(r.FindStringSubmatchIndex("peach punch"))
	fmt.Println(r.FindAllString("peach punch pinch", -1))

	fmt.Println("all:", r.FindAllStringSubmatchIndex("peach punch pinch", -1))
	fmt.Println(r.FindAllString("peach punch pinch", 2)) // limit the matches to 2.
	fmt.Println(r.Match([]byte("peach")))

	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println("regexp:", r)

	// they  expressions can also be used for string modifications.
	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))
}
