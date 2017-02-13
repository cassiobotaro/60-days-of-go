package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println("Strings:")
	// MustCompile parses a regular expression and returns, if successful, a Regexp object that can be used to match against text.
	re := regexp.MustCompile(`p[abc]\d{2,3}(\d{5})`)
	// FindAllString is the 'All' version of FindString; it returns a slice of all successive matches of the expression, as defined by the 'All' description in the package comment. A return value of nil indicates no match.
	for _, e := range re.FindAllString("pb4212345 pa1212345 pc11111111", -1) {
		fmt.Printf("e = %+v\n", e)
	}
	// MatchString reports whether the Regexp matches the string s.
	fmt.Println(re.MatchString("pc0000000"))
	// ReplaceAllString returns a copy of src, replacing matches of the Regexp with the replacement string repl. Inside repl, $ signs are interpreted as in Expand, so for instance $1 represents the text of the first submatch.
	fmt.Println(re.ReplaceAllString("pb12345678another content that lives", ""))

	// Also work with bytes
	fmt.Println("[] byte:")
	for _, e := range re.FindAll([]byte("pb4212345 pa1212345 pc11111111"), -1) {
		fmt.Printf("e = %s\n", e)
	}
	fmt.Println(re.Match([]byte("pc0000000")))
	// intentional not weel formatted to remeber that return are an array of bytes
	fmt.Println(re.ReplaceAll([]byte("pb12345678another content that lives"), []byte("")))

	// for more, read the docs: https://golang.org/pkg/regexp/
}
