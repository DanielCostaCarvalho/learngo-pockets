package main

import (
	"encoding/json"
	"os"
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
