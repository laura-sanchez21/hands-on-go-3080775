// flow-control/loop-while/begin/main.go
package main

import (
	"fmt"
//"strconv"
)

func main() {
	// initialize a variable to check if the loop should continue
	//
	value := 5

	// iterate while the condition is true
	for value >= 0 {

		//fmt.Printf("%v\n", strconv.Itoa(value))
		fmt.Println(value)
		value--
	}
	 
}
