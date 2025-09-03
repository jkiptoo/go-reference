package main

import (
	"fmt"
	"time"
)

//Use Rate Limiting to control resource utilisation and maintenance of quality of service.

func main() {

	//Limit handling of incoming requests by serving the requests off a channel
	incomingRequests := make(chan int, 5)
	for incoming := 1; incoming <=5; incoming++ {
		incomingRequests <- incoming
	}	
	close(incomingRequests)

	//Regulate using rateLimiter to receive a value every 200ms
	rateLimiter := time.Tick(200 * time.Millisecond)

	//Limit to serving 1 request per 200ms and block on receive from the rateLimiter channel
	for request := range incomingRequests {
		<- rateLimiter
		fmt.Println("Batch One Incoming request", request, time.Now())
	}
	
	//Buffer rateLimiter channel  by allowing 3 events in a burst at a time
	limitBursts := make(chan time.Time, 3)

	//Fill up channel for bursting
	for range 3 {
		limitBursts <- time.Now()
	}

	//Add new value every 200ms to maintain limitBursts at the stated capacity of 3 at a time
	go func() {
		for timeLimit := range time.Tick(200 * time.Millisecond) {
			limitBursts <- timeLimit
		}
	}()

	//Simulate 5 additional incoming requests with the first 3 taking advantage of limitBursts

	requestBursts := make(chan int, 5)
	for incoming := 1; incoming <= 5; incoming++ {
		requestBursts <- incoming
	}

	close(requestBursts)
	for request := range requestBursts {
		<- limitBursts
		fmt.Println("Batch Two Incoming request...", request, time.Now())
	}
	//First 3 incoming requests in Batch Two are served immediately unlike in Batch One. The remaining 2 requests are served in 200ms.
}