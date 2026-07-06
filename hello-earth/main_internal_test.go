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

	cases := map[string]testCase{
		"English": {language("en"),"Hello world"},
		"Portuguese": {language("pt"),"Olá, mundo"},
		"Spanish": {language("es"),"Hola, mundo"},
		"Empty string": {language(""),"unsupported language: \"\""},
		"Invalid string": {language("erro"),"unsupported language: \"erro\""},
	}

	for key, currentCase := range cases {
		t.Run(key, func(t *testing.T) {
			language := currentCase.language

			want:= currentCase.want

			got := greet(language)

			if got != want {
				t.Errorf("expected: %q; got: %q", got, want)
			}
		})
	}
}
