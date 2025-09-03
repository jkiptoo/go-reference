package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

//Base management of state is communication over channels e.g. using worker pools.
//Use sync/atomic package to be accessed by multiple goroutines.


func main() {
	//Use atomic integer type for positive counter
	var operation atomic.Uint64
	
	//Use WaitGroup to wait for all goroutines to complete their work.
	var groupWait sync.WaitGroup
	
	//Start 50 goroutines that increment the counter 1000 times each.
	for range 50 {
		groupWait.Add(1)

		go func() {
			for range 1000 {

				//Use .Add method to increment the counter
				operation.Add(1)
			}

			groupWait.Done()
		}()
	}

	//Wait for goroutines to complete
	groupWait.Wait()

	//Use .Load method to safely read values while other goroutines are updating them.
	fmt.Println("Number of Current Operations amounts to a total of...", operation.Load())
}
//Using non-atomic integer would yield different results as the goroutines would run on top of each other. 
//Get data race failures using -race flag