// challenges/types-primitive/begin/main.go
package main

import (
	"fmt"
)

// use type inference to create a package-level variable "val" and assign it a value of "global"
var val string = "temp"

func main() {
	// create a local variable "val" and assign it an int value
	val := 100
	// print the value of the local variable "val"
	fmt.Printf("%v, %T\n",val, val)
	// print the value of the package-level variable "val"
	printingGlobalVariable()
	// update the package-level variable "val" and print it again
	modifyingGlobalVariable("new temp")
	printingGlobalVariable()

	// print the pointer address of the local variable "val"
	//fmt.Printf("%v", &val)
	valAddress := &val
	// assign a value directly to the pointer address of the local variable "val" 
	//*(&val) = 10
	*valAddress = 11
	//and print the value of the local variable "val"
	//fmt.Println("\n", val)
	fmt.Println(val)
}

func printingGlobalVariable() {
	fmt.Println(val)
}

func modifyingGlobalVariable(newString string) {
	val = newString
}
