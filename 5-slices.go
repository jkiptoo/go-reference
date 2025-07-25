package main

import (
	"fmt"
	"slices"
)

func main() {
	
	//Starting Slice Definition/Structure

	var slice []string
	fmt.Println("Un-Initialized Slice:", slice , slice == nil, len(slice) == 0)
	
	// Use 'make' to create a slice with non-zero length
	slice = make([]string, 3)
	fmt.Println("Slice with non zero length:", slice, "Slice Length:", len(slice), "Slice Capacity:", cap(slice))

	//Set and Get just like in Arrays
	slice[0] = "allow"
	slice[1] = "brackets"
	slice[2] = "clap"

	fmt.Println("Set Slice:", slice)
	fmt.Println("Get Slice", slice[2])


	//Return length of slice using Length
	fmt.Println("Length of Slice:", len(slice))


	//Using 'append' method on slices to insert new values
	slice = append(slice, "diagram")
	slice = append(slice, "elephant","flat")
	fmt.Println("Slice with appended values", slice)

	//Copy one slice into another
	sliceCopy := make([]string, len(slice))
	copy(sliceCopy, slice)
	fmt.Println("Display slice copy", sliceCopy)


	//Extract values at certain indices ([2],[3] and [4]) from slices
	extract := slice[2:5]
	fmt.Println("Extracted Slice", extract)


	//Extract slice upto but excluding index 5
	extract = slice[:5]
	fmt.Println("Price with exclusion", extract)


	//Extract slice from and including index [2]
	extract = slice[2:]
	fmt.Println("Slice from specific index", extract)


	//Declare and initialize a variable(s) for slice
	sliceVariableInit := []string{"grand", "halo", "industry"}
	fmt.Println("Declared Variable for Slice", sliceVariableInit)
	
	
	//Slices package contains useful utility functions/methods
	sliceMethods := []string{"grand", "halo", "intel"}
	if slices.Equal(sliceVariableInit, sliceMethods) {
		fmt.Println("sliceVariableInit == sliceMethods")
	}

	//Multi-dimensional Data structures/matrices formed using slices
	twoDimensional := make([][]int, 3)
	for i := range 3 {
		innerLen := i + 1
		twoDimensional[i] = make([]int, innerLen)
		for j := range innerLen {
			twoDimensional[i][j] = i + j
		}
	}
	fmt.Println("2 Dimensional Structure", twoDimensional)
}