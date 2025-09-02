package main

import "fmt"

//Beffered channels accept limited number of values without having to have a corresponding receiver for the values

func main() {
	//Make channely of strings accepting up to 2 values
	messages := make(chan string, 2)
	
	//Send values above into the channel without a receiver
	messages <- "Buffered"
	messages <- "Channel"
	
	//Receive values as usual
	fmt.Println(<-messages)
	fmt.Println(<-messages)
	
}