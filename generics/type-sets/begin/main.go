// generics/type-sets/begin/main.go
package main

import "fmt"

// create a numeric interface with a type set
type numeric interface {
	~float64 | ~int
	grow()
}

// update sum function to use a numeric interface with a type set
func sum[T numeric](a, b T) T {
	return a + b
}

type specialInt int

func(si specialInt) grow() {
	
}

func main() {
	one := specialInt(1)
	two := specialInt(2)
	fmt.Println(sum(one, two))
}
