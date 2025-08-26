package main

import "fmt"

//Function to take many ints as arguments
func sum(numbers ...int) {
	fmt.Print(numbers, "")
	total := 0

	//Type of numbers is equal to []int. Iterate using range. You can also call it using len(nums)
	for _, numbers := range numbers {
		total += numbers
	}
	fmt.Println(total)
}

func main() {

	//Call variadic functions normally using individual arguments
	sum(1,2)
	sum(1,2,3)

	//Apply multiple arguments from a slice into a variadic function e.g. funct(slice)
	numbers := []int{1,2,3,4}
	sum(numbers...)
	
}