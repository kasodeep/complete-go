package handling

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// function to handle the error.
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadingDemo() {
	// directly reading the file.
	dat, err := os.ReadFile("data.txt")
	check(err)
	fmt.Print(string(dat))

	// opening the file to control the operations.
	f, err := os.Open("data.txt")
	check(err)

	// read the first 5 bytes of the files.
	b1 := make([]byte, 5)
	n1, err := f.Read(b1) // n1 denotes the number of bytes read.
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))

	// seek to the 6th byte of the file.
	o2, err := f.Seek(6, io.SeekStart)
	check(err)

	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)

	fmt.Printf("%d bytes @ %d: ", n2, o2)
	fmt.Printf("%v\n", string(b2[:n2]))

	// different functions.
	o3, err := f.Seek(6, io.SeekStart)
	check(err)

	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)

	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	// reset the file pointer.
	_, err = f.Seek(0, io.SeekStart)
	check(err)

	// the bufio maybe useful for the method it provides.
	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))

	f.Close()
}
