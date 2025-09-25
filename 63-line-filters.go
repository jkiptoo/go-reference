//Line filters read input in stdin, process them, then output results to stdout.
//grep and sed are most commonly used

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//This line filter writes capitalised versions of all input text

	//Wrap unbuffered os.Stdin with a buffered scanner
	scanBuffer := bufio.NewScanner(os.Stdin)

	//Scan method takes the scanner to the next token which is in the default scanner 
	for scanBuffer.Scan() {
		
		//Use 'Text' to return the current token which is the next line in the input
		returnCurrentToken := strings.ToUpper(scanBuffer.Text())
		
		//Write the capitalised line
		fmt.Println(returnCurrentToken)	
	}

	//Check for errors. Scan does not report end of file as an error.
	if err := scanBuffer.Err(); 
	err != nil {
		fmt.Fprintln(os.Stderr, "Error Detected:", err)
		os.Exit(1)
	}

}