// Using Go's flag package to specify command line flags like -l to specify options in command line programs
package main

import (
	"flag"
	"fmt"
)

func main() {

	print := fmt.Println

	//Declare a string flag word with a default value of 'list' and a matching description
	//flag.String returns a string pointer, not a string value
	wordFlag := flag.String("word", "list", "a string")

	numberFlag := flag.Int("number", 59, "an integer")
	forkFlag := flag.Bool("fork", false, "a boolean")
	
	//Declare an option that uses a var declared elsewhere in the program then pass a pointer to the flag declaration function
	var redeclareVar string
	flag.StringVar(&redeclareVar, "redeclared variable", "track", "a string var")

	//Call flag.Parse to execute the command-line parsing
	flag.Parse()

	//Dereference the pointers with pointers to get actual option values
	print("Word:", *wordFlag)
	print("Number:", *numberFlag)
	print("Fork:", *forkFlag)
	print("Redeclared var:", redeclareVar)
	print("Trailing Positional Arguments:", flag.Args())
	
}
//----NOTES----//
//Compile then run binary, so as to test the command lines flags program
// Use 'go build command-lines-flags.go'

//Try out the program by giving it values for all programs
// './command-line-flags -word=opt -number=7 -fork -redeclareVar=flag'

//If you remove flags, default values are used
//'./command-line-flags -word=opt'

//Trailing positional arguments can be provided after flags
//'command-line-flags -word=opt a1 a2 a3'

//Use -h or --help flags to get help
//'./command-line-flags -h'

//Display error if a flag is not available in the program
//'./command-line-flags -wat'