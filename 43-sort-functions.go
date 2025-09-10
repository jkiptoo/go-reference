package main

import (
	"cmp"
	"fmt"
	"slices"
)

//Create custom sorts for collections
func main() {
	//Use cmp.Compare comparison function for string lengths
	fruits := []string{"orange", "tamarind", "avocado"}

	compareLength := func(first, second string) int {
		return cmp.Compare(len(first), len(second))
	}
	
	//Call slices.SortFunct to sort fruits by length
	slices.SortFunc(fruits,compareLength)
	fmt.Println(fruits)

	//Use the same method above to sort a slice
	type Person struct {
		name string
		age int
	}
	people := []Person {
		Person{name: "Jackson", age: 55},
		Person{name: "James", age: 26},
		Person{name: "Sven", age: 54},
	}

	//Sort people by age
	slices.SortFunc(people, func(first, second Person) int {
		return cmp.Compare(first.age, second.age)
	})

	fmt.Println(people)

	//If Person struct is large, let slice have *Person instead
}