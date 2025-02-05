// types/conversion/begin/main.go
package main

import "fmt"

func main() {
	// declare float and convert it to an unsigned int
	num := 10.23
	num2 := uint(num)

	fmt.Printf("num: %v=", num)
	fmt.Printf("num2: %v=", num2)

	fmt.Printf("%v=%T, %v=%T", num, num, num2, num2)
}
