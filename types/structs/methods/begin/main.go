// types/structs/methods/begin/main.go
package main

import "fmt"

type author struct {
	first string
	last  string
}

// fullName returns the full name of the author
//This is a method
func (auth author) fullname() string{
	return auth.first + " " + auth.last
}
// This is a function
// func fullName(auth author) {
// 	fmt.Println(auth.first, auth.last)
// }

func main() {
	// initialize author
	a := author {
		first: "james",
		last: "bond",
	}

	// print the author's full name
	//fullName(a)
	fmt.Println(a.fullname())
}
