package main

import "testing"

// Example of how to make tests using table design.

func TestFizz(t *testing.T) {
	// test numbers that are divible by three only.
	// table with input and the expected results.
	table := []struct {
		Input    int
		Expected string
	}{
		{3, "Fizz"},
		{6, "Fizz"},
		{9, "Fizz"},
	}
	// iterate over the table and test if obtained is equals to the expected value
	for _, data := range table {
		if actual := FizzBuzz(data.Input); actual != data.Expected {
			// error will handled if something goes wrong.
			t.Errorf("expected %q but %q was obtained", data.Expected, actual)
		}
	}
}

func TestBuzz(t *testing.T) {
	// test numbers that are divible by five only.
	table := []struct {
		Input    int
		Expected string
	}{
		{5, "Buzz"},
		{10, "Buzz"},
		{20, "Buzz"},
	}
	for _, data := range table {
		if actual := FizzBuzz(data.Input); actual != data.Expected {
			t.Errorf("expected %q but %q was obtained", data.Expected, actual)
		}
	}
}

func TestFizzBuzz(t *testing.T) {
	// test numbers that are divible by three and five.
	table := []struct {
		Input    int
		Expected string
	}{
		{15, "FizzBuzz"},
		{30, "FizzBuzz"},
		{60, "FizzBuzz"},
	}
	for _, data := range table {
		if actual := FizzBuzz(data.Input); actual != data.Expected {
			t.Errorf("expected %q but %q was obtained", data.Expected, actual)
		}
	}
}

func TestNumbers(t *testing.T) {
	// If is not divisible by three or five, return the number as string
	table := []struct {
		Input    int
		Expected string
	}{
		{1, "1"},
		{2, "2"},
		{4, "4"},
	}
	for _, data := range table {
		if actual := FizzBuzz(data.Input); actual != data.Expected {
			t.Errorf("expected %q but %q was obtained", data.Expected, actual)
		}
	}
}
