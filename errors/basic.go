package errors

import (
	"fmt"
	"strconv"
)

// Convention is that the error strings should not capitalize.
func getUser(name string) (string, error) {
	if len(name) > 3 {
		return "", fmt.Errorf("can't find user %s", name)
	}
	return name, nil
}

/*
Go programs express errors with error values.
An Error is any type that implements the simple built-in error interface:
*/
func ErrorDemo() {
	_, err := strconv.Atoi("42A")
	if err != nil {
		fmt.Println("Couldn't convert:", err)
		return
	}

	user, err := getUser("Deep")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(user)
}
