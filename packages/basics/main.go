// packages/basics/main.go
package main

import (
	"fmt"
	"time"
)

func main() {
	const name, age = "kim", 22
	fmt.Println(name, "is", age, "years old")
	fmt.Printf("Today is %s", time.Now().Weekday())

}
