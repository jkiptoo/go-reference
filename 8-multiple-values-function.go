package main

import "fmt"

//Function returning 2 ints
func values()(int, int) {
	return 5, 9
	}

func main() {
// Use of two different return values utilising the function call above	
c, d := values()
fmt.Println(c)
fmt.Println(d)

//Use blank identifier to output subset of returned values
_, e := values()
fmt.Println(e)
}