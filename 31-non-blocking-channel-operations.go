package main

import "fmt"

//Use select with a default case to execute non-blocking sends, receives and multiple selects
func main() {
	
	messages := make(chan string)
	signals := make(chan bool)

	//Non-blocking receive to check if value is present in messages then pass it over to select. 

	select {
	case message := <- messages:
		fmt.Println("Message has been received...", message)
	
	//default case is resorted to if there is no value.	
	default:
		fmt.Println("No message has been received")
	}

	//Message is sent to the messages channel. Channel has no buffer nor receiver hence default case is resorted to.
	message := "Hello there"
	select {
	case messages <- message:
		fmt.Println("Sent message is >>>", message)

	default:
		fmt.Println("No message has been sent.")

	}

	//Many cases are used above the default case to execute a multi-way non-blocking receipt. Trying non-blocking receives on both messages and signals. 
	select {
	case message := <- messages:
		fmt.Println("Message received is...", message)
	
	case signal := <- signals:
		fmt.Println("Received signal is...", signal)
		
	default:
		fmt.Println("No activity detected.")
	}
}