package maps

import (
	"errors"
	"fmt"
)

/*
We created a map of users with their names and phone numbers.
We can create map using make function.
*/
func getUserMap(names []string, phoneNumbers []int) (map[string]user, error) {
	userMap := make(map[string]user)

	if len(names) != len(phoneNumbers) {
		return nil, errors.New("invalid sizes")
	}

	for i := 0; i < len(names); i++ {
		name := names[i]
		phoneNumber := phoneNumbers[i]
		userMap[name] = user{
			name:        name,
			phoneNumber: phoneNumber,
		}
	}
	return userMap, nil
}

type user struct {
	name        string
	phoneNumber int
}

func MapsDemo() {
	names := []string{"John", "Jane"}
	phoneNumbers := []int{123456789, 987654321}

	// getting the user maps.
	userMap, err := getUserMap(names, phoneNumbers)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(userMap)

	// creating the ages map without make.
	ages := map[string]int{
		"John": 37,
		"Mary": 21,
	}
	fmt.Println("Age of Mary:", ages["Mary"])
	delete(ages, "Mary")

	value, ok := ages["Mary"]
	if !ok {
		fmt.Println("key does not exist")
	} else {
		fmt.Println(value)
	}
}
