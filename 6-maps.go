package main

import (
	"fmt"
	"maps"
)

func main() {

	//Create empty map using 'make'
	newMap := make(map[int]string)

	//Set key-value pairs
	newMap[1] = "Antigua"
	newMap[2] = "Bermuda"

	//Print out map with all key-value pairs
	fmt.Println("New Map Countries:", newMap)

	
	//Fetch value of key after passing key parameter
	fetchValue := newMap[1]
	fmt .Println("Fetched Value from key 1 is:", fetchValue)


	//Return number of key-value pairs in a map
	fmt.Println("Map Length", len(newMap)) 


	//Delete key-value pairs from a map
	delete(newMap, 2)
	fmt.Println("Display new map after deletion", newMap)

	
	//Remove all key-value pairs from a map
	clear(newMap)
	fmt.Println("Cleared Map:", newMap)


	//Differentiate between missing keys and keys with zero values using a blank identifier. Only second value is returned
	_, secondReturnValue := newMap[2]
	fmt.Println("The second returned value is:", secondReturnValue)


	//Declare and initialize a new map in one line
	freshMap := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("Newly Declared Map", freshMap)

	
	//Useful Map Functions/Methods
	mapFunction := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(freshMap, mapFunction) {
		fmt.Println("freshMap is equal == mapFunction")
	}
	









}