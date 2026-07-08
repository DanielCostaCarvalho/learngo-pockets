package main

import (
	"flag"
	"fmt"
)

func main() {
	var lang string

	flag.StringVar(
		&lang,
		"lang",
		"pt",
		"The required language, e.g. en, pt, es",
	)
	flag.Parse()

	greet := greet(language(lang))
	fmt.Println(greet)
}

type language string

var phrasebook = map[language]string{
	"en": "Hello world",
	"es": "Hola, mundo",
	"pt": "Olá, mundo",
}

func greet(language language) string {
	phrase, ok := phrasebook[language]

	if !ok {
		return fmt.Sprintf("unsupported language: %q", language)
	}

	return phrase
}
