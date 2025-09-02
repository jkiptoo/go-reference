package main

import "fmt"

func main() {

	//Create a zero-valued array that will hold 5 ints
	var a [5]int
	fmt.Println("Emp:", a)

	//Set value at an Index
	a[4] = 200
	fmt.Println("Set:", a)
	fmt.Println("Get value at Index 4:", a[4])

	
	//Return length of an array using len
	fmt.Println("Len:", len(a))

	
	//Declare and initialise array in one line
	b := [5]int{5,4,3,2,8}
	fmt.Println("Declare:", b)

	
	//Use [...] to count the number of elements in an array
	b = [...]int{5,4,2,3,8,}
	fmt.Println("CountElements:", b)


	//Specify the index for a value using [...]
	b = [...]int{1,3:200,100}
	fmt.Println("SpecifyIndexandValue:", b)

	
	//Use Array types to build multi-dimensional arrays
	var twoD [2][3]int
	for i := range 2 {
		for j := range 3 {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2DimensionalArray:", twoD)

	
	//Create and initialise multi-dimensional arrays at once
	twoD = [2][3]int {
		{2,3,4},
		{4,3,2},
	}
	fmt.Println("2DimensionalArray:", twoD)
}