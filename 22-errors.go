package main

import (
	"errors"
	"fmt"
)

//Go error handling goes hand in hand with functions. Error checking is implemented at every stage of code.

//Errors are of type 'error' and are the last return value
func feedNumbers(arg int) (int, error) {
	if arg == 52 {
		
		//Construct basic error value using 'errors.New' function.
		return -1, errors.New("Can't work with 52")
	}

	//Nil value shows that there is no error present
	return arg + 3, nil
}

//Indicate a specific error condition using sentinel error
var ErrOutOfDrinks = fmt.Errorf("No more drinks available. Please check after few minutes.")
var ErrPower = fmt.Errorf("There is no power therefore the water cannot boil. Please connect and switch on the power.")

func makeDrink(arg int) error {
	if arg == 2 {
		return ErrOutOfDrinks
	} else if arg == 4 {
		//Wrap errors with %w verb to a+show context and can be queried using functions such as errors.Is and errors.As
		return fmt.Errorf("Making drinks: %w", ErrPower) 
	}
	return nil
}

func main() {
	//Use inline error check in the if line
	for _, i := range []int{7, 52} {
		if r, e := feedNumbers(i); e != nil {
			fmt.Println("Feed Numbers failed:", e)
		} else {
			fmt.Println("Feed Numbers worked:", r)
		}
	}

	for i := range 5 {
		if err := makeDrink(i); err != nil {
			//Use error.Is to check that errors or error chains match a specific error value. Useful for wrapped or nested errors allowing indetification of specific error types.
			
			if errors.Is(err, ErrOutOfDrinks) {
				fmt.Println("We should buy new drinks!")
			} else if errors.Is(err, ErrPower) {
				fmt.Println("There is no power.")
			} else {
				fmt.Printf("Unknown Error: %s\n", err)
			}
			continue
		}
		fmt.Println("Drinks are ready.")
	}	
}