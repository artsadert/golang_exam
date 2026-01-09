package task3

import (
	"strings"
	"unicode/utf8"
)

func EncryptWord(word string) string {
	runes := []rune(word)

	for i, j := 1, utf8.RuneCountInString(word)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func EncruptPhrase(phrase string) string {
	words := []string{}

	for _, word := range strings.Split(phrase, " ") {
		words = append(words, EncryptWord(word))
	}

	return strings.Join(words, " ")
}
