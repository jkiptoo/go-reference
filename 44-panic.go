package main

import "os"

//Use panic for errors that we cannot handle
//Use panic to handle errors that are unexpected to occur during an operation
func main() {
	//Use panic to check for unexpected errors
	panic("There is a major problem!")

	//Use panic to abort program if function returns a value that we cannot handle gracefully
	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}	
}