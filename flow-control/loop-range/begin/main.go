// flow-control/loop-range/begin/main.go
package main

import (
	"fmt"
)

func main() {
	// initialize a slice of ints
	numlist := []int{1, 2, 3, 4}
	
	// use for-range to iterate over the slice and print each value
	for i, nums := range numlist {
		fmt.Println("index:", i, " Value:", nums)
	}

	// declare a map of strings to ints
	 stringToInt := map[string]int {
		"one": 1, 
		"two": 2,
		"three": 3,
	}
	
	// use for-range to iterate over the map and print each key/value pair
	for i, value := range stringToInt {
		fmt.Println("index:", i, "Value:", value)
	}
}
