package main

import "testing"

var (
	handmaidsTale = Book{
		Author: "Margaret Atwood", Title: "The Handmaid's Tale",
	}
	oryxAndCrake = Book{
		Author: "Margaret Atwood", Title: "Oryx and Crake",
	}
	theBellJar = Book{
		Author: "Sylvia Plath", Title: "The Bell Jar",
	}
	janeEyre = Book{
		Author: "Charlotte Brontë", Title: "Jane Eyre",
	}
)

func equalBookworms(t *testing.T, bookworms, target []Bookworm) bool {
	t.Helper()

	if len(bookworms) != len(target) {
		return false
	}

	for i := range bookworms {
		if bookworms[i].Name != target[i].Name {
			return false
		}

		if !equalBooks(t, bookworms[i].Books, target[i].Books) {
			return false
		}
	}

	return true
}

func equalBooks(t *testing.T, books, target []Book) bool {
	t.Helper()

	if len(books) != len(target) {
		return false
	}

	for i := range books {
		if books[i].Title != target[i].Title {
			return false
		}

		if books[i].Author != target[i].Author {
			return false
		}
	}

	return true
}

func TestLoadBookworms(t *testing.T) {
	type testCase struct {
		bookwormFile string
		want         []Bookworm
		wantErr      bool
	}

	cases := map[string]testCase{
		"Testdata": {
			"testdata/bookworms.json",
			[]Bookworm{
				{
					Name:  "Fadi",
					Books: []Book{handmaidsTale, theBellJar},
				},
				{
					Name:  "Peggy",
					Books: []Book{oryxAndCrake, handmaidsTale, janeEyre},
				},
			},
			false,
		},
		"Error no file":      {"testdata/inexistent.json", nil, true},
		"Error invalid file": {"testdata/invalid.json", nil, true},
	}

	for key, currentCase := range cases {
		t.Run(key, func(t *testing.T) {
			path := currentCase.bookwormFile

			want := currentCase.want

			got, err := loadBookworms(path)

			if err != nil && !currentCase.wantErr {
				t.Errorf("unexcpected error %q", err.Error())
			}

			if err == nil && currentCase.wantErr {
				t.Errorf("expected an error")
			}

			if !equalBookworms(t, got, want) {
				t.Errorf("expected: %q; got: %q", got, want)
			}
		})
	}
}

func equalBooksCount(t *testing.T, booksCount, target BooksCount) bool {
	t.Helper()

	if len(booksCount) != len(target) {
		return false
	}

	for book, targetCount := range target {
		count, ok := booksCount[book]
		if !ok || targetCount != count {
			return false
		}
	}

	return true
}

func TestBooksCount(t *testing.T) {
	type testCase struct {
		bookworms []Bookworm
		want      BooksCount
	}

	cases := map[string]testCase{
		"No Bookworm": {
			[]Bookworm{},
			BooksCount{},
		},
		"Valid bookworm": {
			[]Bookworm{
				{
					Name:  "Test 1",
					Books: []Book{handmaidsTale, oryxAndCrake},
				},
				{
					Name:  "Test 2",
					Books: []Book{handmaidsTale, theBellJar},
				},
			},
			BooksCount{handmaidsTale: 2, oryxAndCrake: 1, theBellJar: 1},
		},
		"Bookworm with the same book twice": {
			[]Bookworm{
				{
					Name:  "Test 1",
					Books: []Book{handmaidsTale, handmaidsTale, oryxAndCrake},
				},
				{
					Name:  "Test 2",
					Books: []Book{handmaidsTale, theBellJar},
				},
			},
			BooksCount{handmaidsTale: 3, oryxAndCrake: 1, theBellJar: 1},
		},
		"Bookworm without books": {
			[]Bookworm{
				{
					Name:  "Test 1",
					Books: []Book{},
				},
				{
					Name:  "Test 2",
					Books: []Book{},
				},
			},
			BooksCount{},
		},
		"Bookworm with nil books": {
			[]Bookworm{
				{
					Name:  "Test 1",
					Books: nil,
				},
				{
					Name:  "Test 2",
					Books: nil,
				},
			},
			BooksCount{},
		},
	}

	for key, currentCase := range cases {
		t.Run(key, func(t *testing.T) {
			bookworms := currentCase.bookworms

			want := currentCase.want

			got := booksCount(bookworms)

			if !equalBooksCount(t, got, want) {
				t.Errorf("expected: %v; got: %v", got, want)
			}
		})
	}
}

func TestCommonBooks(t *testing.T) {
	type testCase struct {
		bookworms []Bookworm
		want      []Book
	}

	cases := map[string]testCase{
		"No Bookworm": {
			[]Bookworm{},
			[]Book{},
		},
		"Valid bookworm": {
			[]Bookworm{
				{
					Name:  "Test 1",
					Books: []Book{handmaidsTale, oryxAndCrake},
				},
				{
					Name:  "Test 2",
					Books: []Book{handmaidsTale, theBellJar},
				},
			},
			[]Book{handmaidsTale},
		},
		"Bookworm with the same book twice": {
			[]Bookworm{
				{
					Name:  "Test 1",
					Books: []Book{handmaidsTale, handmaidsTale, oryxAndCrake},
				},
				{
					Name:  "Test 2",
					Books: []Book{handmaidsTale, theBellJar},
				},
			},
			[]Book{handmaidsTale},
		},
		"Bookworm with two common books": {
			[]Bookworm{
				{
					Name:  "Test 1",
					Books: []Book{handmaidsTale, oryxAndCrake},
				},
				{
					Name:  "Test 2",
					Books: []Book{handmaidsTale, oryxAndCrake},
				},
			},
			[]Book{oryxAndCrake, handmaidsTale},
		},
		"Bookworm without books": {
			[]Bookworm{
				{
					Name:  "Test 1",
					Books: []Book{},
				},
				{
					Name:  "Test 2",
					Books: []Book{},
				},
			},
			[]Book{},
		},
		"Bookworm with nil books": {
			[]Bookworm{
				{
					Name:  "Test 1",
					Books: nil,
				},
				{
					Name:  "Test 2",
					Books: nil,
				},
			},
			[]Book{},
		},
	}

	for key, currentCase := range cases {
		t.Run(key, func(t *testing.T) {
			bookworms := currentCase.bookworms

			want := currentCase.want

			got := findCommonBooks(bookworms)

			if !equalBooks(t, got, want) {
				t.Errorf("expected: %v; got: %v", got, want)
			}
		})
	}
}
