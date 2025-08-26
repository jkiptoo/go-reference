package main

import "fmt"

//Function that calls itself until it reaches base case of 0
func fact(number int) int {
	if number == 0 {
		return 1
	}
	return number * fact(number-1)
} 

func main() {
	fmt.Println(fact(7))

	//Make anonymous function recursive by declaring a variable using var keyword
	var fibonacci func(number int) int
	fibonacci = func(number int) int {
		if number < 2 {
			return number
		}

		return fibonacci(number-1) + fibonacci(number-2)
	}
	//Fibonacci function is called as it is declared under the function main
	fmt.Println(fibonacci(7))	
}