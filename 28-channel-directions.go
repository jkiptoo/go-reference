package main

import "fmt"

//Use channels as function parameters to specify if a channel is meant to receive or receive values.

//Ping channel only accepts a channel for sending values
func networkSendPing(ping chan <- string, message string) {
	ping <- message
}

//Function with two channels, one for receiving and one for sending
func receiveSendNetwork (ping <- chan string, receive chan <- string) {
	message := <- ping
	receive <- message
}

func main() {
	ping := make(chan string, 1)
	receive :=make(chan string, 1)
	networkSendPing(ping, "Message has been passed!")
	receiveSendNetwork(ping, receive)
	fmt.Println(<- receive)
}