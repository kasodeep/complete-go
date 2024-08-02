package main

import "fmt"

func main() {
	// type conversion.
	accountAgeFloat := 2.6
	accountAgeInt := int32(accountAgeFloat)
	fmt.Println("Your account has existed for", accountAgeInt, "years")
}
