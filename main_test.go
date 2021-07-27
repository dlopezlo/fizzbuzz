package main

import "testing"

func TestFizzBuzz(t *testing.T) {
	tt := []struct {
		number uint64
		result string
	}{
		{1, "1"},
		{2, "2"},
		{3, "Fizz"},
		{4, "4"},
		{5, "Buzz"},
		{9, "Fizz"},
		{10, "Buzz"},
		{15, "FizzBuzz"},
		{30, "FizzBuzz"},
		{50, "Buzz"},
		{55, "Buzz"},
		{500, "Buzz"},
		{933, "Fizz"},
	}

	for _, tc := range tt {
		r := fibo(tc.number)
		if r != tc.result {
			t.Errorf("fibo(%v) got %v and should be %v", tc.number, r, tc.result)
		}
	}
}
