// Some string formatting tasks in go lang
package main

import (
	"fmt"
	"os"
)

//Create struct for axis points
type axis struct {
	x, y int16
}

func main() {
	//Print instance of axis struct
	a := axis{5, 7}
	fmt.Printf("First axis points to: %v\n", a)

	//Use %+v verb to include struct's field names
	fmt.Printf("Second axis points to: %+v\n", a)

	//Print a Go syntax representing the value using %#v verb
	fmt.Printf("Third axis points to: %#v\n", a)

	//To print type of value, use %T verb
	fmt.Printf("Type of axis: %T\n", a)

	//Standard base10 formatting uses %d verb
	fmt.Printf("Standard base 10 formating: %d\n", 125)

	//Use %t verb to format booleans
	fmt.Printf("Boolean formatting: %t\n", true)

	//Binary representation uses the %b verb
	fmt.Printf("Binary representation: %b\n", 17)

	//Use %c verb to print character that corresponds to integer value
	fmt.Printf("Character representation is: %c\n", 32)

	//For hex encoding, use %x
	fmt.Printf("Hex encoding is %x\n", 400)

	//For basic decimal formatting, use %f verb
	fmt.Printf("Basic float is: %f\n", 80.3)

	//Use %e and %E to format floats in scientific scenarios 
	fmt.Printf("Scientific float representation is: %e\n", 15904011.0)
	fmt.Printf("Scientific float notation part two: %E\n", 15904011.0)

	//Use %s verb for basic string printing
	fmt.Printf("String one: %s\n", "\"STRING\"")

	//Use %q verb to double quote strings
	fmt.Printf("String two: %q\n", "\"STRING\"")

	//To get two characters per byte of input, use %x verb which provides base16 representation of a string
	fmt.Printf("String three: %x\n", "Represent Hex")

	//Use %p verb for pointer formatting
	fmt.Printf("Pointer is represented as: %p\n", &a)

	//Formatting width of numbers to be right-justified and padded with spaces
	fmt.Printf("The width for normal numbers is: |%6d|%6d|\n", 15, 400)

	//Specify width of floats using width.precision convention
	fmt.Printf("The width for floating numbers is: |%6.2f|%6.2f|\n", 50.2, 100.2)

	//To justify number on the left, use - flag
	fmt.Printf("To left justify floating numbers: |%-6.2f|%-6.2f|\n", 50.2, 100.2)

	//Output strings and format them to be right-justified and ensure they are shown in a table-like format
	fmt.Printf("Formatting strings and aligning them in a table: |%6s|%6s|\n", "flat", "beta")

	//Left justify strings
	fmt.Printf("Left justify strings: |%-6s|%-6s|\n", "flat", "beta")

	//Sprintf formats and returns a string without printing the said string anywhere
	//Sprintf is mainly used with os.Stdout
	string := fmt.Sprintf("Sprintf in action: a %s", "STRING")
	fmt.Println(string)


	//Using Fprintf to format and print to io.Writers instead of os.Stdout
	fmt.Fprintf(os.Stderr, "Input and Output: an %s\n", "Error detected")
	
}