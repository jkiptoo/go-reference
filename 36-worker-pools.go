package main

import (
	"fmt"
	"time"
)

//Using goroutines and channels to implement a worker pool

//In this worker, several concurrent instances will be run.
//Work is received on the 'currentJob' channell and output is sent to 'jobResults' channel.
func worker(id int, currentJob <- chan int, jobResults chan <- int) {
	for executingJob := range currentJob {
		fmt.Println("Worker", id, "Worker has started the job.", executingJob)
		time.Sleep(time.Second)
		fmt.Println("Worker", id, "Worker has finished the job", executingJob)
		jobResults <- executingJob * 2
	}
}

func main() {
	//Make 2 channels for sending jobs to workers and collecting their results
	const numberOfJobs = 5
	currentJob := make(chan int, numberOfJobs)
	jobResults := make(chan int, numberOfJobs)

	//Start 3 initially blocked channels
	for processWorker := 1; processWorker <= 3; processWorker++ {
		go worker(processWorker, currentJob, jobResults)
	}

	//Send 5 jobs then close the channel to complete the work
	for secondProcessWorker := 1; secondProcessWorker <= numberOfJobs; secondProcessWorker++ {
		currentJob <- secondProcessWorker
	}
	close(currentJob)

	//Collect results of work done. Alternative is to use 'WaitGroup'
	for workerResults := 1; workerResults <= numberOfJobs; workerResults++ {
		<- jobResults
	}

}
//5 jobs are executed by workers.
//Program lasts for two seconds because 3 workers are operating concurrently.
//Program was to last 5 seconds total.