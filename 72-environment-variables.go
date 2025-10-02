// Using environmental variables to pass configuration information in Unix programs
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	print := fmt.Println

	//If key is not present in the environment, empty string is returned. 
	//Use os.Setenv to set a key/value pair.
	//Use os.Getenv to get value of a key
	os.Setenv("SystemConfig", "1")
	print("SystemConfig:", os.Getenv("SystemConfig"))
	print("SystemSetup", os.Getenv("SystemSetup"))

	print()  //Print all keys
	
	//List all key/value pairs in the environment using os.Environ
	for _, environment := range os.Environ() {

		//Return a slice of strings in the form of 'KEY=value'
		pair := strings.SplitN(environment, "=", 2)
		print(pair[0])

	}
}