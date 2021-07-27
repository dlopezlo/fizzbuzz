package main

import "testing"

func TestFizzBuzz(t *testing.T) {
	tt := []struct {
		name   string
		number uint64
		result string
	}{
		{"a number should return the same number", 1, "1"},
		{"3 should return Fizz", 3, "Fizz"},
		{"5 should return Buzz", 5, "Buzz"},
		{"a multiple of 3 should return Fizz", 9, "Fizz"},
		{"a multiple of 5 should return Buzz", 10, "Buzz"},
		{"A multiple of both 3 and five should return FizzBuzz", 15, "FizzBuzz"},
	}

	for _, tc := range tt {
		r := fibo(tc.number)
		if r != tc.result {
			t.Errorf("[Test name: %v] fibo(%v) got %v and should be %v", tc.name, tc.number, r, tc.result)

		}
	}
}
