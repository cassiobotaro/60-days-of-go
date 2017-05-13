package main

import (
	"fmt"
	"math"
	"sort"
	"sync"
)

import "C"

// extracted from https://dev.to/vladimirvivien/calling-go-functions-from-other-languages

var count int
var mtx sync.Mutex

//export add
func add(a, b int) int {
	return a + b
}

//export cosine
func cosine(x float64) float64 {
	return math.Cos(x)
}

//export sortints
func sortints(vals []int) {
	sort.Ints(vals)
}

//export logmsg
func logmsg(msg string) int {
	mtx.Lock()
	defer mtx.Unlock()
	fmt.Println(msg)
	count++
	return count
}

func main() {}
