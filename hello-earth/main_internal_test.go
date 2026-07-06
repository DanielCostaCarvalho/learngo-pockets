package main

import "testing"

func Example_main()  {
	main()
	// Output:
	// Olá, mundo
}

func TestGreet(t *testing.T)  {
	type testCase struct {
		language language
		want string
	}

	cases := [...]testCase{
		{language("en"),"Hello world"},
		{language("pt"),"Olá, mundo"},
		{language("es"),"Hola, mundo"},
		{language(""),"unsupported language: \"\""},
		{language("erro"),"unsupported language: \"erro\""},
	}

	for i := range cases {
		language := cases[i].language

		want:= cases[i].want

		got := greet(language)

		if got != want {
			t.Errorf("expected: %q; got: %q", got, want)
		}
	}
}
