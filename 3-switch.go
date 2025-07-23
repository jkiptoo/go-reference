package main

import (
	"fmt"
	"time"
)

func main() {

	//Basic Switch
	i := 3
	fmt.Println("Write", i, " as ")
	switch i {
		case 1: 
		fmt.Println("One")
		
		case 2: 
		fmt.Println("Two")

		case 3: 
		fmt.Println("Three")
	}

	//Use of commas in Switch Statements
	switch time.Now().Weekday() {
		case time.Saturday, time.Sunday:
		fmt.Println("It is the weekend!")

		default:
			fmt.Println("It is a weekday")
	}

	//Switch withouh an expression. Express If/Else Logic.
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It is before noon.")
	default:
		fmt.Println("It is after noon.")
	}

	//Switch statement for comparing types instead of values
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
			case bool:
				fmt.Println("I am a boolean.")
			case int:
				fmt.Println("I am an integer")
			default:
				fmt.Printf("I don't know my type %T\n", t)
		}
		
	}
	
	
	//Finally get output of type passed into the "whatAmI" function
	whatAmI(true)
	whatAmI(5)
	whatAmI(4.32)
	whatAmI("Hey you there")

}