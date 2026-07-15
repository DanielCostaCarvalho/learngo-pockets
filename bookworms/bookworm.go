package main

import (
	"bufio"
	"encoding/json"
	"os"
)

type Bookworm struct {
	Name  string `json:"name"`
	Books []Book `json:"books"`
}

// lê o arquivo e faz parse para slice de Bookworm
func loadBookworms(filepath string) ([]Bookworm, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	buffedReader := bufio.NewReaderSize(file, 1024*1024)

	var bookworms []Bookworm

	err = json.NewDecoder(buffedReader).Decode(&bookworms)
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
