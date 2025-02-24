// flow-control/loop-basic/begin/main.go
package main

import (
	"fmt" 
)
func main() {
	// declare a string to iterate over
	//stringToIterate := "Hello World"
	stringToIterate := "Hello"
	// iterate over the string with basic for loop
	for i := 0; i < len(stringToIterate); i++ {
		//fmt.Printf("%c\n", stringToIterate[i])
		fmt.Println(i, ":", string(stringToIterate[i]))
	}
}
