package arrays

import "fmt"

func ArrayDemo() {
	var myInts [10]int
	myInts[0] = 1
	fmt.Println("MyInts:", myInts, "Length:", len(myInts))

	primes := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("Primes:", primes, "Length:", len(primes))

	// slice operation
	// slice are references and hence changes will be reflected in the original array.
	fmt.Print("Slice Operations: ")
	fmt.Println(primes[1:4], primes[1:], primes[:4], primes[:])

	// make function.
	// func make([]T, len, cap) []T
	mySlice := make([]int, 4, 10)
	mySlice = append(mySlice, 1, 2, 3, 4)
	fmt.Println("Slice using make(...):", mySlice)

	// address.
	fmt.Println("Primes address:", &primes[0])
}
