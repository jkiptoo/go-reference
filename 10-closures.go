package main

// Closures are formed using anonymous functions which help define a function inline without having to name it

import "fmt"

//Function 'integerSequence' that returns another function defined anonymously within intSequence

func integerSequence() func() int {
	i := 0

	//Close function over variable i to form a closure
	return func() int {
		i++
		return i
	} 
}

func main() {
	//Call integerSequence and assign function result to nextInteger. The value of i gets update everytime nextInteger is called
	nextInteger := integerSequence()

	//Call closure several times to see effect of closure
	fmt.Println(nextInteger())
	fmt.Println(nextInteger())
	fmt.Println(nextInteger())

	newInteger := integerSequence() //Confirm uniqueness of state to that particular function by testing new function
	fmt.Println(newInteger())

}