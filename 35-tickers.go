package main

import (
	"fmt"
	"time"
)

//Tickers are executing processes at regular intervals.

func main() {
	//Create channel to wait for arival of values every 500ms
	intervalTicker := time.NewTicker(500 * time.Millisecond)
	tickerComplete := make(chan bool)

	go func() {
		for {
			select {
			case <- tickerComplete:
				return

			case timing := <- intervalTicker.C:
				fmt.Println("Ticker is at an interval of...", timing)

			}
		}
	}()
	
	//Cancel and stop ticker so as to not receive anymore values from the channel. 
	time.Sleep(1600 * time.Millisecond)
	intervalTicker.Stop()
	tickerComplete <- true
	fmt.Println("Ticker has been successfully stopped.")
}
//Program ticks 3 times before it is stopped.