package structs

import "fmt"

/* Anonymous struct are used when you want to create a struct without a name.
These structs can only be used once.
*/
func Anonymous() {
	myCar := struct {
		make  string
		model string
	}{
		make:  "Ford",
		model: "Mutsang",
	}

	fmt.Println(myCar)
}
