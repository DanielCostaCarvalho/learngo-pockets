package main

import (
	"encoding/json"
	"os"
	"sort"
)

type Bookworm struct {
	Name  string `json:"name"`
	Books []Book `json:"books"`
}

type Book struct {
	Author string `json:"author"`
	Title  string `json:"title"`
}

// lê o arquivo e faz parse para slice de Bookworm
func loadBookworms(filepath string) ([]Bookworm, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var bookworms []Bookworm

	err = json.NewDecoder(file).Decode(&bookworms)
	if err != nil {
		return nil, err
	}

	return bookworms, nil
}

type BooksCount map[Book]uint

// retorna a contagem de livros
func booksCount(bookworms []Bookworm) BooksCount {
	count := make(BooksCount)

	for _, bookworm := range bookworms {
		for _, book := range bookworm.Books {
			count[book]++
		}
	}

	return count
}

// retorna livros em comum
func findCommonBooks(bookworms []Bookworm) []Book {
	booksCount := booksCount(bookworms)

	var commonBooks []Book

	for book, count := range booksCount {
		if count > 1 {
			commonBooks = append(commonBooks, book)
		}
	}

	return sortBooks(commonBooks)
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
