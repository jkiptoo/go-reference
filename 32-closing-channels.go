package main

import "fmt"

//Signals completion to the channel's receivers
//No more values will be sent to a channel

func main() {
	//Create a channel named 'processJobs' that will indicate the work to be done by the main() goroutine.

	processJobs := make(chan int, 5)
	jobComplete := make(chan bool)

	//Worker goroutine that receives from 'processJobs'. If all values have been received, notifications is sent to 'jobDone' once all jobs have been worked on.

	go func() {
		for {
			currentJob, moreJobs := <- processJobs
			if moreJobs {
				fmt.Println("Received job is...", currentJob)
			} else {
				fmt.Println("All jobs have been received and processed.")
				jobComplete <- true
				return
			}
		}
	} ()

	//Send 3 jobs to the worker 'go func' over the 'processJobs' channel
	for currentJob := 1; currentJob <= 3; currentJob++ {
		processJobs <- currentJob
		fmt.Println("Job has been sent", currentJob)
	}

	//Close the 'processJobs' channel
	close(processJobs)
	fmt.Println("All jobs have been sent and processed.")

	//Await worker using synchronization approach
	<- jobComplete
	
	//Read from a closed channel if the value received is a zero value of the type. Return value is true if a send operation delivered a value. Return value is false if a zero value is generated because the channel is closed and empty.

	_, affirmative := <- processJobs
	fmt.Println("Have more jobs been received?", affirmative)


}