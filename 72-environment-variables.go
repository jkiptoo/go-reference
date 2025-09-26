// Using environmental variables to pass configuration information in Unix programs
package main

import (
	"fmt"
	"os"
)

func main() {
	print := fmt.Println

	//If key is not present in the environment, empty string is returned. 
	//Use os.Setenv to set a key/value pair.
	//Use os.Getenv to get value of a key
	os.Setenv("SystemConfig")

}