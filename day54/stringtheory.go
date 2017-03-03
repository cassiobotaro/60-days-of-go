package main

import "fmt"

// Square ...
type Square struct {
	side int
}

// String method is needed to implements Stringer interface.
// Stringer is implemented by any value that has a String method,
// which defines the “native” format for that value.
// The String method is used to print values passed as an operand to any format
// that accepts a string or to an unformatted printer such as Print.
func (s Square) String() string {
	return fmt.Sprintf("square : %d", s.side)
}

// GoString method is needed to implements GOStringer interface.
// GoStringer is implemented by any value that has a GoString method,
// which defines the Go syntax for that value.
// The GoString method is used to print values passed as an operand to a %#v format.
func (s Square) GoString() string {
	return fmt.Sprintf("square side: %d area: %d", s.side, s.side*s.side)
}

// Rectangle ...
type Rectangle struct {
	height int
	width  int
}

func main() {
	square := Square{3}
	// call method String
	fmt.Println(square)
	fmt.Printf("%v\n", square)
	// call method GoString
	fmt.Printf("%#v\n", square)

	rectangle := Rectangle{1, 2}
	// without name in fields
	fmt.Printf("%v\n", rectangle)
	// with name in fields
	fmt.Printf("%+v\n", rectangle)
	// more details about the struct
	fmt.Printf("%#v\n", rectangle)

	// Type
	fmt.Printf("%T\n", 10)

	// binary
	fmt.Printf("%b\n", 10)

	// unicode
	fmt.Printf("%c\n", 666)
	fmt.Printf("%c\n", 42)
	fmt.Printf("%c\n", 1024)

	fmt.Printf("base 10: %d\nbase 8: %o\nbase 16 lower-case: %x\nbase16 upper-case: %X\n", 42, 42, 42, 42)
	// %q	a single-quoted character literal safely escaped with Go syntax.
	fmt.Printf("%q\n", "text with quotes")
	// %U	Unicode format: U+1234; same as "U+%04X"
	fmt.Printf("%U\n", 42)

	// boolean
	fmt.Printf("%t\n", true)
	// chan
	c := make(chan int)
	fmt.Printf("%p\n", c)
	// pointer
	fmt.Printf("%p\n", &rectangle)
	// left justify
	fmt.Printf("|%-6s|%-6s|\n", "hello", "world")
	// right align
	fmt.Printf("|%6s|%6s|\n", "hello", "world")
	// float precision
	fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)
	// to see more about string formatting:
	// - https://gobyexample.com/string-formatting
	// - https://golang.org/pkg/fmt

}
