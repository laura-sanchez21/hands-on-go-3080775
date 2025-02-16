// types/maps/initialization/begin/main.go
package main

import "fmt"

type author struct {
	name string
}

func main() {
	// declare a map of string keys and author values
	//var authors map[string]author
	authors := map[string]author {
		"tm": {name: "toni morrison"},
		"ma": {name: "Marcus Aurelius"},
	}

	// initialize the map with make
	//authors = make(map[string]author)

	// add authors to the map
	//authors["ma"] = author{name: "Marcus Aurelius"}
	//authors["tm"] = author{name: "Toni Morrison"}

	// print the map with fmt.Printf
	//fmt.Printf("%#v\n", authors)
	
	// read a value from the map with a known key
	//fmt.Printf("%#v\n", authors["ma"])
	fmt.Println(authors["tm"])
}
