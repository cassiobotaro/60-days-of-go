package main

import (
	"math/rand"
	"reflect"
	"testing"
	"testing/quick"
)

// Generate is the needed method to implements generate interface
// needed to generate random data
func (r Retangle) Generate(rand *rand.Rand, size int) reflect.Value {
	randomRetangle := Retangle{
		height: rand.Int(),
		width:  rand.Int(),
	}
	return reflect.ValueOf(randomRetangle)

}

// Test table like style
func TestIsOddTable(t *testing.T) {
	// tipical golang test
	// anonymous struct to iterate over cases
	for _, value := range []struct {
		expected bool
		input    int
	}{
		{true, 3},
		{false, 2},
		{true, -5},
		{false, -4},
	} {
		if obtained := IsOdd(value.input); value.expected != obtained {
			t.Errorf("Expected %t but found %t", value.expected, obtained)
		}
	}
}

// Tests uding testing/quick package
func TestNumberPlusOne(t *testing.T) {
	// assertion: if n is odd -> n+1 is even and opposite is truth too
	assertion := func(x int) bool {
		// you will see this log many times
		t.Log("Random value: ", x)
		return IsOdd(x) != IsOdd(x+1)
	}
	// check the assertion for many random cases(default 100)
	if err := quick.Check(assertion, nil); err != nil {
		t.Error(err)
	}
}

// As Retangle implents generate, we can use your own function to generate random data
func TestRetangle(t *testing.T) {
	assertion := func(x Retangle) bool {
		// again we see many iterations
		t.Log("Random Retangle:", x)
		return (x.height == x.width) && (x.area() == x.height*x.height) || x.height != x.width
	}
	if err := quick.Check(assertion, nil); err != nil {
		t.Error(err)
	}
}
