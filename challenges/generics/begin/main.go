// challenges/generics/begin/main.go
package main

import (
"fmt" 
"golang.org/x/exp/constraints"
)

// Part 1: print function refactoring
func printAny[T any](output T){
	fmt.Println(output)
}
// non-generic print functions

func printString(s string) { fmt.Println(s) }

func printInt(i int) { fmt.Println(i) }

func printBool(b bool) { fmt.Println(b) }

// Part 2 sum function refactoring
type numeric interface {
	constraints.Float | constraints.Integer
}

func sumAny[T numeric](nums ...T) T {
	var sum T
	for _, num := range nums {
		sum += num
	}
	return sum
}

// sum sums a slice of any type
func sum(numbers []interface{}) interface{} {
	var result float64
	for _, n := range numbers {
		switch n.(type) {
		case int:
			result += float64(n.(int))
		case float32, float64:
			result += n.(float64)
		default:
			continue
		}
	}
	return result
}

func main() {
	// call non-generic print functions
	// printString("Hello")
	// printInt(42)
	// printBool(true)

	// call generic printAny function for each value above
	printAny("Hello")
	printAny(42)
	printAny(true)

	// call sum function
	//fmt.Println("result", sum([]interface{}{1, 2.1, 3}))

	// call generics sumAny function
	//list := []int {1, 2, 3}
	fmt.Println("result", sumAny(1, 2.9, 3))
}
