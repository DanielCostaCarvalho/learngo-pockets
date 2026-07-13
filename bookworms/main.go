package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var filepath string

	flag.StringVar(
		&filepath,
		"path",
		"",
		"The path to the json file with the bookworms",
	)
	flag.Parse()

	bookworms, err := loadBookworms(filepath)
	if err != nil {
		fmt.Printf("Failed to load bookworms: %s\n", err)
		os.Exit(1)
	}

	commonBooks := findCommonBooks(bookworms)

	fmt.Println("Here are the books in common:")

	for _, book := range commonBooks {
		fmt.Println(book.toString())
	}
}
