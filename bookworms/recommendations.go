package main

func recommendOtherBooks(bookworms []Bookworm) []Bookworm {
	bookRecommendations := make(bookRecommendations)

	for _, bookworm := range bookworms {
		for i, book := range bookworm.Books {
			otherBooksOnShelves := listOtherBooksOnShelves(i, bookworm.Books)
			registerBookRecommendations(bookRecommendations, book, otherBooksOnShelves)
		}
	}

	recommendations := make([]Bookworm, len(bookworms))

	for i, bookworm := range bookworms {
		recommendations[i] = Bookworm{
			Name:  bookworm.Name,
			Books: recommendBooks(bookRecommendations, bookworm.Books),
		}
	}

	return recommendations
}

type bookRecommendations map[Book]bookCollection
type bookCollection map[Book]struct{}

func (allBooks bookCollection) Contains(book Book) bool {
	_, ok := allBooks[book]
	return ok
}

func (allBooks bookCollection) toListOfBooks() []Book {
	list := make([]Book, 0, len(allBooks))

	for book := range allBooks {
		list = append(list, book)
	}

	return sortBooks(list)
}

func listOtherBooksOnShelves(currentIndex int, books []Book) []Book {
	otherBooks := make([]Book, 0, len(books)-1)

	for i, book := range books {
		if i != currentIndex {
			otherBooks = append(otherBooks, book)
		}
	}

	return sortBooks(otherBooks)
}

func registerBookRecommendations(bookRecommendations bookRecommendations, book Book, otherBooksOnShelves []Book) {
	for _, otherBook := range otherBooksOnShelves {
		recommendationsMap, ok := bookRecommendations[book]
		if !ok {
			recommendationsMap = make(bookCollection)
			bookRecommendations[book] = recommendationsMap
		}

		recommendationsMap[otherBook] = struct{}{}
	}
}

func recommendBooks(bookRecommendations bookRecommendations, myBooks []Book) []Book {
	allRecommendations := make(bookCollection)

	myBooksCollection := make(bookCollection)

	for _, book := range myBooks {
		myBooksCollection[book] = struct{}{}
	}

	for _, book := range myBooks {
		for recommendation := range bookRecommendations[book] {
			if myBooksCollection.Contains(recommendation) {
				continue
			}

			allRecommendations[recommendation] = struct{}{}
		}
	}

	return allRecommendations.toListOfBooks()
}
