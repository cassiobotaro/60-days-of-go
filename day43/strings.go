package main

import (
	"fmt"
	"strings"
)

func main() {
	// Fun with strings functions
	// I choose 5 functions that looks funny

	// Fields splits the string s around each instance of one or more consecutive white space characters, as defined by unicode.IsSpace, returning an array of substrings of s or an empty list if s contains only white space.
	for _, word := range strings.Fields("a \n string\t with spaces") {
		fmt.Printf("word = %+q\n", word)
	}

	// IndexAny returns the index of the first instance of any Unicode code point from chars in s, or -1 if no Unicode code point from chars is present in s.
	// first vowel index
	i := strings.IndexAny("some random content", "aeiou")
	fmt.Printf("i = %+v\n", i)

	// Join concatenates the elements of a to create a single string. The separator string sep is placed between elements in the resulting string.
	phrase := strings.Join([]string{"an", "array", "of", "words"}, "  ")
	fmt.Printf("phrase = %+v\n", phrase)

	// Map returns a copy of the string s with all its characters modified according to the mapping function. If mapping returns a negative value, the character is dropped from the string with no replacement.
	newString := strings.Map(func(original rune) rune {
		return original + 300
	}, "random string")
	fmt.Printf("newString = %+v\n", newString)

	// TrimSpace returns a slice of the string s, with all leading and trailing white space removed, as defined by Unicode.
	fmt.Printf("%q\n", strings.TrimSpace(" \t\t  text with trash \n\t \r    \n"))

	// for more, read the docs: https://golang.org/pkg/strings
}
