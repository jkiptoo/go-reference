package main

import (
	"fmt"
	"time"
)

//Use timeouts in programs that connect to external resources

func main() {
	//Prevent goroutine leaks in unread channel.
	//The sending go routine is buffered so not blocked
	//Execute an external call that returns result to channelOne after 2 seconds.

	channelOne := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		channelOne <- "Result one is now out..."
	}()

	//result waits for the result
	select {
	case result := <- channelOne:
		fmt.Println(result)
	
	//wait for value sent after timeout of 1 second. Timeout case is activated if execution takes more than 
	case <- time.After(1 * time.Second):
		fmt.Println("Timeout of 1 second")	
	}

	//Timeout longer than 3 seconds leads to receipt from channelTwo and output is displayed
	channelTwo := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		channelTwo <- "Result two is out..."
	} ()

	//Output of results
	select {
	case result := <- channelTwo:
		fmt.Println(result)
	
	case <- time.After(3 * time.Second):
		fmt.Println("Timeout of more than 3 seconds")	
	}
}