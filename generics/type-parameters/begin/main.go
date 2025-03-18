// generics/type-parameters/begin/main.go
package main

import "fmt"

func sumInts(a, b int) int {
	return a + b
}

func sumFloats(a, b float64) float64 {
	return a + b
}

// type Number interface {
// 	int | float64
// }

// create generic sum function with type parameter T constrained to int and float64 types
// func sum[T Number](nums []T) T{
// 	var result T
// 	for _,num := range nums {
// 		result += num
// 	}
// 	return result
// }

//~ allows for support on custom types
func sum[T ~int | ~float64](a, b T) T {
	return a + b
}

type specialInt int

func main() {
	// non-generic sum int function
	fmt.Println(sumInts(1, 2))

	// non-generic sum float function
	fmt.Println(sumFloats(1.3, 2.2))

	// call on generic sum function
	fmt.Println("Sum floats:", sum(1.2, 1.3))
	fmt.Println("Sum ints:", sum(4, 5))

	// define a compatible custom type call on generic sum function with it
	one := specialInt(1)
	two := specialInt(2)
	fmt.Println(sum(one, two))

	temp := &list[int] {next: nil, val: 7}
	temp1 := &list[int] {next: temp, val: 8}

	fmt.Println(temp1)
	
}

// list is a singly-linked list that holds values of any type
type list[T any] struct {
	next *list[T]
	val T
}
