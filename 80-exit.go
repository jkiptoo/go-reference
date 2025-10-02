// Using 'os.Exit' to exit with a given status.
package main

import (
	"fmt"
	"os"
)

func main() {
	print := fmt.Println
	//When using 'os.Exit', 'defer' will not run and 'print' will not be called.
	defer print("What is happening!!!?")

	//Exit with status 3
	os.Exit(3)

}

//-------------NOTES-------------------------//
//Go does not use an integer return value from main to indicate exit status
//To exit with  non-zero status, use 'os.Exit'
//Status is shown in the terminal after building and executing a binary.