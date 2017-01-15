package main

// Decorator receives a pipe function and modify your behavior
func Decorator(f func(s string) string) func(s string) string {
	return func(s string) string {
		// change received parameter adding more text
		return f(s + " World!")
	}
}

func main() {
	// simple pipe function that receives a string and return itself
	pipe := func(s string) string {
		return s
	}
	// decorate pipe
	decorated := Decorator(pipe)
	println(decorated("Hello"))
}
