package main

import "fmt"

func main() {
	var x uint64
	for x = 1; x <= 45; x++ {
		fmt.Println(fibo(x))
	}
}

func fibo(number uint64) (result string) {
	if number%3 == 0 && number%5 == 0 {
		return "FizzBuzz"
	}
	if number%3 == 0 {
		return "Fizz"
	}
	if number%5 == 0 {
		return "Buzz"
	}
	return fmt.Sprintf("%v", number)
}
