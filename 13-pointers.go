package main

import "fmt"

//Use pointers to pass references to values. zeroValue function has an int parameter and will get a copy of ivalue.
func zeroValue(ivalue int) {
	ivalue = 0
}

//Passes Pointer using *. '*iparameter' code dereferences the pointer from its memory address to the current value of the address 
func zeroParameter(iparameter *int) {
	*iparameter = 0
} 


func main() {

	i := 1
	fmt.Println("Initial value is:", i)
	
	//zeroValue does not change the value of i in func main. zeroParameter does change the value because it has reference to the memory address
	zeroValue(i)
	fmt.Println("Zero Value is:", i)

	//Use &i syntax to give the memory address of i
	zeroParameter(&i)
	fmt.Println("Zero Parameter is:", i)

	//Print out pointer memory address
	fmt.Println("Pointer memory address is:", &i)
	
}