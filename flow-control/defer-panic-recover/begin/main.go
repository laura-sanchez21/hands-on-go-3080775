// flow-control/defer-panic-recover/begin/main.go
package main

import "fmt"

func cleanup(msg string) {
	fmt.Println(msg)
}

func main() {
	// defer function calls
	//panic("something bad happened")
	defer cleanup("first")
	defer cleanup("second")

	fmt.Println("working in main...")
	// defer recovery
	defer func(){
		if err := recover(); err != nil {
			fmt.Println("recovered from panic:", err)
		}
	}()
	// panic
	panic("something bad happened")
}

// func safeDivide(a, b int) {
// 	// Recovering from panic
// 	defer func() {
// 		fmt.Println("defer func")
// 		if r := recover(); r != nil {
// 			fmt.Println("Recovered from error:", r)
// 		}
// 	}()

// 	if b == 0 {
// 		fmt.Println("b == 0\n")
// 		panic("Cannot divide by zero!") // Triggers panic
// 	}
// 	fmt.Println("Result:", a/b)
// }

// func main() {
// 	//safeDivide(10, 2) // Works fine
// 	safeDivide(10, 0) // Panics but gets recovered
// 	fmt.Println("Program continues running...")
// }