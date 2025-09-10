package main

import (
	"fmt"
	"slices"
)

//Sorting for built in types
func main() {
	//Sorting functions work for any built-in type
	sortStrings := []string{"cal", "aby", "blk"}
	slices.Sort(sortStrings)
	fmt.Println("Sort strings:", sortStrings)

	//Sort ints
	sortIntegers := []int{8,4,2}
	slices.Sort(sortIntegers)
	fmt.Println("Sort integers:", sortIntegers)

	//Check if slice is already sorted
	 sortedSlice := slices.IsSorted(sortIntegers) 
	 fmt.Println("Sorted slice: ", sortedSlice)	
}