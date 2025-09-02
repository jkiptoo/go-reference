package main

import (
	"fmt"
	"sync"
	"time"
)

//Use WaitGroups when awaiting multiple goroutines to finish. When passing WaitGroup to functions, use pointers

//Create function to run in every goroutine
func workerProcess(jobId int) {
	fmt.Printf("Worker %d is starting\n", jobId)

	//Using sleep to demonstrate an expensive task
	time.Sleep(time.Second) 
	fmt.Printf("Worker %d is done processing\n", jobId)
}


func main() {
	//Using WaitGroup to wait for all goroutines to finish
	var processWait sync.WaitGroup

	//Start several goroutines using WaitGroup.Go
	for processRoutine := 1; processRoutine <= 5; processRoutine++ {
		processWait.Go(func() {
			workerProcess(processRoutine)
		})
	}

	//Block until all the goroutines are completed
	processWait.Wait()
}