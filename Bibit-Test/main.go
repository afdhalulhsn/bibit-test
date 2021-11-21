package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	words := []string{"kita", "atik", "tika", "aku", "kia", "makan", "kua"}
	result := Anagram(words)
	bracket := "Afdhal (Ganteng)"

	fmt.Println(result)
	fmt.Println("Word In Bracket=", findFirstStringInBracket(bracket))
}

//Anagram
func Anagram(words []string) [][]string {
	var anagramGroup [][]string
	mapSortWord := map[string][]string{}
	for _, w := range words {
		chars := strings.Split(w, "")
		sort.Strings(chars)
		sortKey := strings.Join(chars, "")
		mapSortWord[sortKey] = append(mapSortWord[sortKey], w)
	}
	for _, maps := range mapSortWord {
		anagramGroup = append(anagramGroup, maps)
	}
	return anagramGroup
}

//Refactor Code
func findFirstStringInBracket(str string) string {
	if len(str) > 0 {
		if indexFirstBracketFound := strings.Index(str, "("); indexFirstBracketFound >= 0 {
			wordsAfterFirstBracket := str[indexFirstBracketFound:]
			if indexClosingBracket := strings.Index(wordsAfterFirstBracket, ")"); indexClosingBracket >= 0 {
				return wordsAfterFirstBracket[1:indexClosingBracket]
			}
		}
	}
	return ""
}
