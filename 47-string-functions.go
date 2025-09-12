package main

import (
	"fmt"
	s "strings"
)

//Using standard library 'strings' to access useful string functions

var print = fmt.Println


func main() {
	//Example usage of some of the functions available in the 'strings' package
	
	print("String contains: ", s.Contains("First String Test", "es"))
	
	print("Count characters in a string: ", s.Count("Second String Test", "T"))
	
	print("Check if string has prefix: ", s.HasPrefix("Third String Test", "te"))
	
	print("Check if string has suffix: ", s.HasSuffix("Fourth String Test", "st"))
	
	print("Check string index: ", s.Index("Fifth String Test", "e"))
	
	print("Check join index: ", s.Join([]string{"Alto", "Beta"}, "-"))

	print("Repeat string: ", s.Repeat("a", 5))

	print("Replace string: ", s.Replace("foo", "o", "0", -1))

	print("Replace string again: ", s.Replace("foo", "o", "0", 1))

	print("Split string: ", s.Split("a-b-c-d-e", "-"))

	print("Make string lowercase: ", s.ToLower("STING"))

	print("Make string uppercase: ", s.ToUpper("sting"))

}