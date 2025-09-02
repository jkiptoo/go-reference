package main

import "fmt"

//Use channels to connect concurrent goroutines.
//Send values from one goroutine into a channel then receive it in different goroutine.

func main() {
	//Create new channel using make(chan value-type)
	message := make(chan string)

	//Send value into channel using <- syntax. Send ping into message channel.
	go func() { message <- "Ping origin." } ()

	//<- is now used to receive values from the channel above. Now receive message from above channel.
	newMessage := <- message
	fmt.Println(newMessage)
}