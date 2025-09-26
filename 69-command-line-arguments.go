// Using parameterise execution of programs
package main

import (
	"fmt"
	"os"
)

func main() {

	print := fmt.Println

	//Access raw command-line arguments.
	//First slice value gives path to program.
	//os.Args[1:] contains arguments to the program.
	argumentsWithProgram := os.Args
	argumentsWithoutProgram := os.Args[1:]

	//Get individual arguments with normal indexing
	individualArguments := os.Args[3]

	print(argumentsWithProgram)
	print(argumentsWithoutProgram)
	print(individualArguments)
}
// Command 'go build 69-command-line-arguments.go'
//Build a Go binary using 'go build' first 
// Run './69-command-line-arguments a b c d'