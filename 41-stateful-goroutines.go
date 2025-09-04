package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

//Using goroutines and channel approach by having each data piece owned by/dedicated to 1 goroutine

//To guarantee incorruptible data access, read or write state by sending and receiving messages on the dedicated goroutine.

//Using readOperation and writeOperation to encapsulate requests.

type readOperation struct {
	key int
	response chan int
}

type writeOperation struct {
	key int
	value int
	response chan bool
}

func main() {
	//Count the number of operations to perform
	var readOperations uint64
	var writeOperations uint64

	//Create read and write channels to be used to read and write requests.
	reads := make(chan readOperation)
	writes := make(chan writeOperation)

	//Use map goroutine to own state and make the goroutine private
	go func() {
		var state = make(map[int]int)

		//In repetition, select reads and writes channels and repond to requests as they arrive
		for {
			select {
				//Perform requested operation to execute a response
				case read := <- reads:
					read.response <- state[read.key]
				
				//Send a value to response channel to indicate success
				case write := <- writes:
					state[write.key] = write.value
					write.response <- true
			}
		}
	}()

	//Start 500 goroutines
	for range 500 {
		//Use reads channel to read the state-owning goroutine
		go func() {
			for {
				//Construct readOperation and send each read over the reads channel
				read := readOperation {
					key: rand.Intn(5),
					response: make(chan int)}

					//Receive response and output via the response channel
					reads <- read
					<- read.response
					atomic.AddUint64(&readOperations, 1)
					time.Sleep(time.Millisecond)
			}
		}()
	}
	
	//Start 20 writes using the same approach
	for range 20 {
		go func() {
			for {
				write := writeOperation {
					key: rand.Intn(5),
					value: rand.Intn(500),
					response: make(chan bool)}
					writes <- write
					<- write.response
					atomic.AddUint64(&writeOperations, 1)
					time.Sleep(time.Millisecond)
			}
		}()
	}

	//Allow goroutines to work for some time.
	time.Sleep(time.Second)

	//Record and report operation counts.
	readOperationsResult := atomic.LoadUint64(&readOperations)
	fmt.Println("Result of Read Operations is...", readOperationsResult)
	
	writeOperationsResult := atomic.LoadUint64(&writeOperations)
	fmt.Println("Result of Write Operations is...", writeOperationsResult)
}

//SUMMARY: Goroutine state management approach is more involving than using mutexes. Useful where other managing multiple mutexes will cause errors. Also useful in handling cases where other channels are present.