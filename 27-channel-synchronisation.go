package main

import (
	"fmt"
	"time"
)

//Use channels to synchronise excution of goroutines. Use WaitGroup when awaiting the excution of multiple goroutines.

//Done channel is used to notify another goroutine that the function's work is done.
func workerProcess(processDone chan bool) {
		fmt.Print("Worker process is working...")
		time.Sleep(time.Second)
		fmt.Println("Process is done!")

		//Send a value to notify completion of process
		processDone <- true 
	}

func main() {
	//Start worker gorutine and feed it with the channel to send the notification
	processDone := make(chan bool, 1)
	go workerProcess(processDone)

	//Block until notification from the worker is received in the channel
	
	<- processDone

}