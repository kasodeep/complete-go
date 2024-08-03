/*
	int  int8  int16  int32  int64
	uint uint8 uint16 uint32 uint64 uintptr
	float32 float64
	complex64 complex128

	The number represents the number of bits.
	The default type of int is 32 bits and uint is 64 bits.
*/

package variables

import "fmt"

// Capital letter meaning public and small meaning private to current package.
func Types() {
	// declaring.
	var smsSendingLimit int
	var costPerSMS float64
	var hasPermission bool
	var username string
	fmt.Printf("Default Values: %v %.2f %v %q\n", smsSendingLimit, costPerSMS, hasPermission, username)
}
