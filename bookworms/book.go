package main

import (
	"fmt"
	"sort"
)

type Book struct {
	Author string `json:"author"`
	Title  string `json:"title"`
}

// retorna livros ordenados
func sortBooks(books []Book) []Book {
	sort.Slice(books, func(index1, index2 int) bool {
		if books[index1].Author != books[index2].Author {
			return books[index1].Author < books[index2].Author
		}

		return books[index1].Title < books[index2].Title
	})
	return books
}

func (book Book) toString() string {
	return fmt.Sprintf("- %s by %s", book.Title, book.Author)
}
