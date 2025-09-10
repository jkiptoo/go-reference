package main

import (
	"fmt"
)

//Use recover to stop program from aborting when a panic error occurs.

//Panic Function
func abortProgram() {
	panic("There is a major problem! Program will now abort...")

}

func main() {
	//Call recover only inside a defer function to catch the panic call.
	//Recover takes in the value output by the panic error function above
	defer func() {
		if recoverProgram := recover(); recoverProgram !=nil {
		fmt.Println("Program has recovered from a Panic Error:\n", recoverProgram)
	}
}()
abortProgram()

//The code below does not run because the main function stops at the point of the panic then starts again in the deferred closure
fmt.Println("Result after abortProgram() Function")
}
