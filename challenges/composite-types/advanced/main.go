package main

import (
	"fmt"
)

type author struct {
	name string
}

type book struct {
	title string
	author 
}

type library struct {
	books []book
}

func (lib *library) AddBook(b book) {
	lib.books = append(lib.books, b)
}

func (lib library) LookupByAuthorName(author string) (books []book) {
	for _, b := range lib.books { 
        	if(b.author.name == author) {
			books = append(books, b)
		} 
    	}
	return
}

func main() {
	lib := library{
		[]book{},
	}

	lib.AddBook(book{
		title: "Intro to GoLang",
		author: author{
			name: "Author Name 1",
		},
	})

	lib.AddBook(book{
		title: "Advanced Concepts in GoLang",
		author: author{
			name: "Author Name 2",
		},
	})

	fmt.Println(lib)

	booksByAuthor := lib.LookupByAuthorName("Author Name 2")

	for _, b := range booksByAuthor {
		fmt.Println(b)
	}

}
