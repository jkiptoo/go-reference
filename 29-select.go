package main

import (
	"fmt"
	"time"
)

//Use select to combine goroutines and channels and wait on multiple channel operations.

func main() {
	
	//Select from channels
	firstChannel := make(chan string)
	secondChannel := make(chan string)

	//Each channel will receive a value after a period of time e.g. blocking RPC operations from operating in concurrent goroutines.

	go func() {
		time.Sleep(1 * time.Second)
		firstChannel <- "Channel One."
	}()

	go func() {
		time.Sleep(2 * time.Second)
		secondChannel <- "Channel Two.."
	}()

	//Use select to await the two values printing each as they arrive.
	for range 2 {
		select {
		case messageOne := <- firstChannel:
			fmt.Println("Received first message!", messageOne)

		case messageTwo := <- secondChannel:
			fmt.Println("Received second message!", messageTwo)	
		}
	}	
}