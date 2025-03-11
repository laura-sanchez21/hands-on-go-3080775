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
	//floatNums := []float64{1.2, 1.3}
	//intNums := []int{2, 3}
	//fmt.Println("float sum: ", sum(floatNums))
	//fmt.Println("int nums: ", sum(intNums))
	fmt.Println("Sum:", sum[float64](1.4, 2.7))
	fmt.Println("Sum:", sum[int](3, 5))
	//[int | float64] not needed. These can be inferred to by the compile through function
	// argument type inference. When including them it is known as instantiation of 
	// the generic function 

	// define a compatible custom type call on generic sum function with it
	one := specialInt(1)
	two := specialInt(2)
	fmt.Println("Special Case:", sum(one,two))
}

// list is a singly-linked list that holds values of any type
type list[T any] struct {
	next *list[T]
	val T
}
