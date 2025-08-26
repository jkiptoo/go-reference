package main

import "fmt"

//Function that takes two ints and returns their sum
func addition(a int, b int) int {
	return a + b
}

//Omit the type after each parameter in case of many parameters of the same type

func consecutiveAddition(a,b,c,d int) int {
	return a + b + c + d
}


//Now call function using the convention name(args)
func main() {
	result := addition(3,5)
	fmt.Println("3 plus 5 equals to", result)

	result = consecutiveAddition(2,4,6,8)
	fmt.Println("2 plus 4 plus 6 plus 8 equals to", result)
}