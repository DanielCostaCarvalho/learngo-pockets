package main

import "testing"

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
