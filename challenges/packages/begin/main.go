// challenges/packages/begin/proverbs.go
package main

// import the proverbs package
import (
	"fmt"

	"github.com/jboursiquot/go-proverbs"
)

// getRandomProverb returns a random proverb from the proverbs package
func getRandomeProverb() string {
	return proverbs.Random().Saying
}

func main() {
	// print the result of calling your getRandomProverb function
	fmt.Println(getRandomeProverb())
}
