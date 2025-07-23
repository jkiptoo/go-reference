package main

import "fmt"

func main() {
	// Basic example
	if 7%2 == 0 {
	fmt.Println("Number is even")
	} else {
		fmt.Println("Number is odd")
	}
	
	//If statement without an else statement 
	if 8%4 == 0 {
		fmt.Println("Number is divisible by 4")
	}

	//Use of Logical operators such as Logical OR || and Logical AND &&
	if 8%2 == 0 || 7%2 == 0 {
		fmt.Println("Either 8 or 7 are even numbers.")
	}

	// If, Else If, and Else Statement Flow
	if num:= 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has one digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}