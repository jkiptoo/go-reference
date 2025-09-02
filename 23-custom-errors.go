package main

import (
	"errors"
	"fmt"
)

//Use of custom types as errors

//A custom error type has 'Error' appended to it
type argumentError struct {
	argument int
	message string
}

//Error method causes argumentError to implement error interface
func (errorInterface *argumentError) Error() string {
	return fmt.Sprintf("%d - %s", errorInterface.argument, errorInterface.message)
	
}

func age(arg int) (int, error) {
	if arg == 52 {
		return -1, &argumentError{arg, "We cannot work with it."} //Return custom error
	}
	return arg + 3, nil
}

func main() {
	//Use errors.As is a more advanced version of errors.Is. Below errors.As matches specific error type and converts to a value. If no match, false is returned.
	_, err := age(52)
	var errorAs *argumentError
	if errors.As(err, &errorAs) {
		fmt.Println(errorAs.argument)
		fmt.Println(errorAs.message)
	} else {
		fmt.Println("Error does not match argumentError.")
	}

	
}