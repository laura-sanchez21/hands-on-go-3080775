// types/structs/embed/begin/main.go
package main

import "fmt"

type person struct {
	first string
	last  string
}

// fullName returns the full name of a person
func (p person) fullName() string {
	return p.first + " " + p.last
}

// define author and embed person
//
type author struct {
	person person
	penName string
}

// override fullName method for author
//
func (a author) fullName() string {
	return fmt.Sprintf("%s (%s)", a.person.fullName(), a.penName)
}

func main() {
	// initialize and print a person's full name
	p := person{
		first: "Toni",
		last:  "Morrison",
	}
	a := author {
		person: person {
			first: "James",
			last: "Baldwin",
		},
		//person: p,
		penName: "Jimmy",
	}
	fmt.Println(p.fullName())

	// initialize and print an author's full name
	//
	fmt.Println(a.fullName())
}
