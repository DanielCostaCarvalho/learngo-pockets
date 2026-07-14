package main

import "testing"

var (
	book1 = Book{
		Author: "Author 1", Title: "Book 1",
	}
	book2 = Book{
		Author: "Author 2", Title: "Book 2",
	}
	book3 = Book{
		Author: "Author 3", Title: "Book 3",
	}
	book4 = Book{
		Author: "Author 4", Title: "Book 4",
	}
	book5 = Book{
		Author: "Author 5", Title: "Book 5",
	}
)

func TestRecommendOtherBooks(t *testing.T) {
	type testCase struct {
		bookworms []Bookworm
		want      []Bookworm
	}

	cases := map[string]testCase{
		"No common books": {
			bookworms: []Bookworm{
				{Name: "test", Books: []Book{book1, book2}},
				{Name: "test2", Books: []Book{book3, book4}},
			},
			want: []Bookworm{
				{Name: "test"},
				{Name: "test2"},
			},
		},
		"Only common books": {
			bookworms: []Bookworm{
				{Name: "test", Books: []Book{book1, book2}},
				{Name: "test2", Books: []Book{book1, book2}},
			},
			want: []Bookworm{
				{Name: "test"},
				{Name: "test2"},
			},
		},
		"Common books with recommendations": {
			bookworms: []Bookworm{
				{Name: "test", Books: []Book{book1, book2}},
				{Name: "test2", Books: []Book{book2, book3}},
			},
			want: []Bookworm{
				{Name: "test", Books: []Book{book3}},
				{Name: "test2", Books: []Book{book1}},
			},
		},
	}

	for key, currentCase := range cases {
		t.Run(key, func(t *testing.T) {
			want := currentCase.want

			got := recommendOtherBooks(currentCase.bookworms)

			if !equalBookworms(t, got, want) {
				t.Errorf("expected: %q; got: %q", got, want)
			}
		})
	}
}
