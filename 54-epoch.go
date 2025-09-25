// Common requirement in programs is to extract time in number of seconds, milliseconds, or nanoseconds since Unix epoch

package main

import (
	"fmt"
	"time"
)

func main() {
	//Using 'time.Now' to get elapsed time
	print := fmt.Println 
	currentTime := time.Now()
	fmt.Println(currentTime)

	print(currentTime.Unix())
	print(currentTime.UnixMilli())
	print(currentTime.UnixNano())

	//Convert integer seconds or nanoseconds that have elapsed since the Unix epoch
	print(time.Unix(currentTime.Unix(), 0))
	print(time.Unix(0, currentTime.UnixNano()))
}