// Using Go lang's inbuilt time and duration support
package main

import (
	"fmt"
	"time"
)

func main() {
	//Get current time
	print := fmt.Println
	currentTime := time.Now()
	print(currentTime)

	//Build time struct for a specific timezone
	formerTime := time.Date(2024, 1, 1, 12, 0, 0, 0,  time.UTC)	
	print(formerTime)

	//Get various components of time value
	print(formerTime.Year())
	print(formerTime.Month())
	print(formerTime.Day())
	print(formerTime.Hour())
	print(formerTime.Minute())
	print(formerTime.Second())
	print(formerTime.Nanosecond())
	print(formerTime.Local().Location())

	//Output weekday format
	print(formerTime.Weekday())

	//Compare currentTime and formerTime
	print(formerTime.Before(currentTime))
	print(formerTime.After(currentTime))
	print(formerTime.Equal(currentTime))

	//Return duration between currentTime and formerTime using 'Sub' methods
	durationDifference := currentTime.Sub(formerTime)
	print(durationDifference)

	//Compute different duration lengths 
	print(durationDifference.Hours())
	print(durationDifference.Minutes())
	print(durationDifference.Seconds())
	print(durationDifference.Nanoseconds())

	//Using 'Add' or '-' to increase/decrease time by a given duration.
	print(formerTime.Add(durationDifference))
	print(formerTime.Add(-durationDifference))

}