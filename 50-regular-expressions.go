package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {
	//Directly create pattern to match a string
	matchString, _ := regexp.MatchString("o([a-z]+)ge", "orange")
	fmt.Println(matchString)

	//Use Compile to create an optimised Regex struct
	regexTask, _ := regexp.Compile("o([a-z]+)ch")

	//From the methods available to the structs, do a match test
	fmt.Println(regexTask.MatchString("orange"))

	//Find the match for regexp
	fmt.Println(regexTask.FindString("orange change"))
	
	//Find the first match and return the start and end indexes for the match instead of the matching text
	fmt.Println("Index:", regexTask.FindStringIndex("orange change"))

	//Use submatch to include information about whole-pattern matches and submatches within
	fmt.Println(regexTask.FindStringSubmatch("orange change"))
	
	//Return the indexes of matches and submatches
	fmt.Println(regexTask.FindStringSubmatchIndex("orange change"))

	//Use All variants to apply to all matches in the input 
	fmt.Println(regexTask.FindAllString("orange change range", -1))

	//All variants are available in other functions as well
	fmt.Println("All variants:", regexTask.FindAllStringSubmatchIndex("orange change range", -1))

	//Using a non-negative integer to reduce the number of matches
	fmt.Println(regexTask.FindAllString("orange change range", 2))

	//Provide []byte arguments and stop using String
	fmt.Println(regexTask.Match([]byte("orange")))

	//For global variables, use 'MustCompile' instead of 'Compile' so as to return panic instead of an error
	regexTask = regexp.MustCompile("o([a-z]+)ch")
	fmt.Println("Global Regex Expression is: ", regexTask)

	//Uuse regex to replace subset of strings with other values
	fmt.Println(regexTask.ReplaceAllString("an orange", "<fruit>"))

	//Use Func to transform matched text with a function
	inRegex := []byte("an orange")
	outRegex := regexTask.ReplaceAllFunc(inRegex, bytes.ToUpper)
	fmt.Println(string(outRegex))
	
}