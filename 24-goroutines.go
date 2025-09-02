package main

import (
	"fmt"
	"time"
)

func functionCall(from string) {
	for i := range 3 {
		fmt.Println(from, ":", i)
	}
}

func main() {
	//Make function call running synchronously
	functionCall("Direct")
	
	//Create another go routine to execute concurrently with the function above
	go functionCall("GoRoutine") 

	//Start goroutine for anonymous function call
	go func(message string) {
		fmt.Println(message)
	} ("Process is Going.")

	//Two function calls running in an asynchronous manner
	time.Sleep(time.Second)
	fmt.Println("Process Done.")

	//Output is the blocking call first then the output of the two goroutines follows.
}