package main

import (
	"fmt"
)


func main() {
// Basic Loop Start
i := 1
for i <= 3 {
	fmt.Println(i)
	i = i + 1
}
//Basic Loop End

//Initialiser/Condition/After 'FOR LOOP'
for j := 0; j < 3; j++ {
	fmt.Println(j)
}

//Using for range
for k := range 3 {
	fmt.Println("range", k)
}

//Loop repeatedly until you break out of loop
for {
	fmt.Println("loop")
	break
}


//Using for range to continue to the next iteration of loop
for n:= range 6 {
	if n % 2 == 0 {
		continue
	}
	fmt.Println(n)
}
}