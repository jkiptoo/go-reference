// Using pattern-based layouts for time formatting and parsing
package main

import (
	"fmt"
	"time"
)

func main() {
	//Using layout constant to format time in observance of RFC3339
	print := fmt.Println

	currentTime := time.Now()
	print(currentTime.Format(time.RFC3339))

	//Use same layout values as 'Format' when Time parsing
	parseCurrentTime, e := time.Parse( time.RFC3339, "2022-12-02T22:19:41+00:00")
	print(parseCurrentTime)

	//Using 'Format' and 'Parse' as example based layouts. Custom layouts can be used as well.
	print(currentTime.Format("10:01PM"))
	print(currentTime.Format("Wed Sep _3 22:01:05 2022"))
	print(currentTime.Format("2022-09-03T22:01:05.999999-07:00"))
	timeFormat := "10 04 PM"
	parseTime, e := time.Parse(timeFormat, "11 30 PM")
	print(parseTime)

	//Use standard string formatting for the time formatting and parsing output in the above sections.
	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n", currentTime.Year(), currentTime.Month(), currentTime.Day(), currentTime.Hour(), currentTime.Minute(), currentTime.Second())

	//Return error on malformed input giving more detail to the time parsing problem.
	malformedError := "Wed Sep _3 22:01:05 2022"
	_, e = time.Parse(malformedError, "11 30PM")
	print(e)
}