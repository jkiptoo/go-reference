// Handling Unix signals like:
// Shutting down server when it receives 'SIGTERM'
// Stop processing command line input if 'SIGINT' signal is received.
// Signals are handled in Go using channels.
package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	print := fmt.Println()
	//Create a buffered channel to receive notifications.
	//Send 'os.Signal' values into the channel.
	signalNotification := make(chan os.Signal, 1)
	
	//Register the 'signalNotification' channel to receive notifications of signals using 'signal.Notify'
	signal.Notify(signalNotification, syscall.SIGINT,syscall.SIGTERM)

	//Receive 'signalNotification' in separate goroutine instead of the main function.
	done := make(chan bool, 1)

	//Create a blocking receive goroutine.
	//Goroutine receives signal, prints it out, then notifies the program that it can finish.
	go func ()  {

		signalReceipt := <- signalNotification
		print()	
		print(signalReceipt)
		done <- true
		
	}()

	//Program will now wait for signal by sending a value on 'done'
	print("Waiting for signal to arrive...")
	<-done
	print("Exiting Program...")
	
}

//------NOTES------//
//When run, the program will blocking wait for signal.
//Typing 'CTRL + C' sends SIGINT signal.
//Program prints interrupt then exits.