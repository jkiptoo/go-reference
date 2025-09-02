package main

import "fmt"

//Use range to iterate over values received in a channel

func main() {
	//Iterate over 'messageQueue' channel
	messageQueue := make(chan string, 2)
	messageQueue <- "First in queue."
	messageQueue <- "Second in queue."

	close(messageQueue)

	//Using range to iterate over each value received from the messageQueue.
	//Have remaining values received even when channel is closed.

	for elements := range messageQueue {
		fmt.Println(elements)
	}	
}