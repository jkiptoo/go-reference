package main

import (
	"fmt"
	"os"
)

//Use defer to ensure function call is done for cleanup purposes and is always executed at the last stages of a program's execution
func main() {
	//Main idea: Create a file, write to it, then close it.
	
	//Get file object with createFile
	fileObject := createFile("/tmp/defer.txt")

	//Defer the closing of the file using closeFile
	defer closeFile(fileObject)

	//Execute at the end of main function using writeFile function call
	writeFile(fileObject)
}

//Function to create file
func createFile(creatingFile string) *os.File {
	fmt.Println("Creating File...")
	fileObject, err := os.Create(creatingFile)
	if err != nil {
		panic(err)
	}
	return fileObject 
}

//Function to write the file
func writeFile(fileObject *os.File) {
	fmt.Println("Writing File...")
	fmt.Fprintln(fileObject, "Here is the data...")
}

//Function to close the file
func closeFile(fileObject *os.File) {
	fmt.Println("We are now closing the file...")

	//Check for errors when file is closing
	err := fileObject.Close()
	if err != nil {
		panic(err)
	}
}