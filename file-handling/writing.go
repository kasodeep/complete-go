package handling

import (
	"bufio"
	"fmt"
	"os"
)

func WritingDemo() {
	// writing to a file.
	d1 := []byte("hello\ngo\n")
	err := os.WriteFile("data.txt", d1, 0644)
	check(err)

	// creating a new file.
	f, err := os.Create("text.txt")
	check(err)

	defer f.Close()

	// writing to the file.
	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	n3, err := f.WriteString("writes\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n3)

	f.Sync()

	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n4)

	w.Flush() // to ensure all operations executed.
}
