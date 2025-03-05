// challenges/interfaces/begin/main.go
package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"

	"github.com/davecgh/go-spew/spew"
)

type counter interface {
	name() string
	count(input string) int
}

type letterCounter struct{ identifier string }
func (lC letterCounter) name() string {return lC.identifier }
func (lc letterCounter) count(input string) int {
	var count = 0
	for _,c := range input {
		if unicode.IsLetter(c) {
			count++
		}
	} 
	return count
}

type numberCounter struct{ designation string }
func (nC numberCounter) name() string {return nC.designation}
func (nC numberCounter) count(input string) int {
	var count = 0
	for _,c := range input {
		if unicode.IsNumber(c) {
			count++
		}
	}
	return count
}

type symbolCounter struct{ label string }
func (sC symbolCounter) name() string {return sC.label}
func (sC symbolCounter) count(input string) int {
	var count =  0
	for _,c := range input {
		if !unicode.IsLetter(c) && !unicode.IsNumber(c) {
			count++
		}
	}
	return count
}


func doAnalysis(data string, counters ...counter) map[string]int {
	// initialize a map to store the counts
	analysis := make(map[string]int)

	// capture the length of the words in the data
	analysis["words"] = len(strings.Fields(data))

	// loop over the counters and use their name as the key
	for _, c := range counters {
		analysis[c.name()] = c.count(data)
	}

	// return the map
	return analysis
}

func main() {
	// handle any panics that might occur anywhere in this function
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Error:", r)
		}
	}()

	// use os package to read the file specified as a command line argument
	bs, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(fmt.Errorf("failed to read file: %s", err))
	}

	// convert the bytes to a string
	data := string(bs)
	spew.Dump(data)

	// call doAnalysis and pass in the data and the counters
	numCounter := numberCounter{designation: "Number"}
	letterCounter := letterCounter{identifier: "Letter"}
	symbolCounter := symbolCounter{label: "Symbol"}

	analysis := doAnalysis(data, numCounter, letterCounter, symbolCounter)
	// dump the map to the console using the spew package
	 spew.Dump(analysis)
}
