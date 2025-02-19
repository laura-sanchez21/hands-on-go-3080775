// challenges/types-composite/begin/main.go
package main

import (
	//"fmt"

	"fmt"

	"github.com/davecgh/go-spew/spew"
)

// define an author type with a name
//
type author struct {
	name string
}

// define a book type with a title and author
type book struct {
	title string
	author author
}

// define a library type to track a list of books
type library map[string][]book

// define addBook to add a book to the library
func (lib library) addBook(bookToAdd book) {
	//lib.bookList = append(lib.bookList, bookToAdd)
	lib[bookToAdd.author.name] = append(lib[bookToAdd.author.name], bookToAdd)
}

// define a lookupByAuthor function to find books by an author's name
//
func (lib library) lookupByAuthorName(name string) []book {
	//var books []book
	// for i := 0; i < len(lib.bookList); i++ {
	// 	if lib.bookList[i].author.name == author.name {
	// 		books = append(books, lib.bookList[i])
	// 	}
	// }
	//return books

	return lib[name]
}

func main() {
	// create a new library
	 lib := library{}

	// add 2 books to the library by the same author
	jb := author{name: "James Baldwin"}
	// var book1 = book{
	// 	title: "The Fire Next Time",
	// 	author: jb,
	// }

	// var book2 = book{ 
	// 	title: "Black Swan",
	// 	author: jb,
	// }

	//lib.bookList[book1.author.name] = append(lib.bookList[book1.author.name], book1)
	//lib.bookList[book2.author.name] = append(lib.bookList[book2.author.name], book2)
	//lib.addBook(book1)
	//lib.addBook(book2)
	lib.addBook(book{
		title: "The Fire next Time",
		author: jb,
	})

	// add 1 book to the library by a different author
	// var book3 = book {
	// 	title: "Eclipse",
	// 	author: author{name: "Got7"},
	// }
	lib.addBook(book{
		title: "Go Tell It is on the Mountain",
		author: jb,
	})

	lib.addBook(book{
		title: "Meditations",
		author: author{name: "Marcus Aurelius"},
	})

	// dump the library with spew
	spew.Dump(lib)
	//fmt.Printf("%v",lib)

	// lookup books by known author in the library
	books := lib.lookupByAuthorName(jb.name)
	spew.Dump(books)

	// print out the first book's title and its author's name
	if len(books) != 0 {
		b := books[0]
		fmt.Println(b.title, "by", b.author.name)
	}

}
