package main

import "fmt"

//Use range to iterate over elements in various data structures

func main() {

	//Use range to sum numbers in a slice
	numbers := []int{100,200,300, 5}
	sum := 0
	
	//Use blank identifier to ignore index and only publish value
	for _, numbers := range numbers {
		sum += numbers
	}
	fmt.Println("Sum of Slice is:", sum)


	//Iterate using range and publish both the index and value
	for i, numbers := range numbers {
		if numbers == 300 {
			fmt.Println("Index of 300 is:", i)
		}
	}

	//Use range to iterate over key/value pairs in a map
	keyValuePair := map[string]string{"a": "Apricot", "b": "Banana"}
	for key, value := range keyValuePair {
		fmt.Printf("%s -> %s\n", key, value)
	}

	//Using range to iterate over only keys of a map
	for key := range keyValuePair {
		fmt.Println("Key is:", key)
	}

	//Using range on strings iterates over Unicode code
	for i, c := range "Go" {
		fmt.Println(i, c)
	}	
}