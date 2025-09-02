package main

import (
	"fmt"
	"time"
)

//Timers are relevant for executing processes later in the future.

func main() {
	//Timer to reperesent event in the future e.g. waiting for 2 seconds in the future
	firstTimer := time.NewTimer(2 * time.Second)

	//<- firstTimer.C block on the timer channel until it sends a value indicating the timer executed
	<-firstTimer.C
	fmt.Println("The first Timer has finished executing.")

	//Cancel timer before it executes
	secondTimer := time.NewTimer(time.Second)
	go func() {
		<- secondTimer.C
		fmt.Println("The second Timer has finished executing.")
	}()

	stopSecondTimer := secondTimer.Stop()
	if stopSecondTimer {
		fmt.Println("The second time has been cancelled and stopped.")
	}

	//Give secondTimer time to fire then get stopped. Using time.Sleep is important when you want to wait
	time.Sleep(2 * time.Second)
}
//First timer executes 2 seconds after program is started