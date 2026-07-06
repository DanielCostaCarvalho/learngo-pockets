package main

import "fmt"

func main()  {
	greet := greet("pt")
	fmt.Println(greet)
}

type language string

var phrasebook = map[language]string{
	"en":   "Hello world",
	"es":   "Hola, mundo",
	"pt":  "Olá, mundo",
}

func greet(language language) string {
	phrase, ok := phrasebook[language]

	if !ok {
		return fmt.Sprintf("unsupported language: %q", language)
	}

	return phrase
}
