package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	const helloString = "こんにちは"

	//Strings are equivalent to byte. Length of the raw bytes is produced in the code below
	fmt.Println("Length of String is:", len(helloString))

	
	//Loop to generate hex values of all the value that make up the code points of helloString

	for i := 0; i < len(helloString); i++ {
		fmt.Printf("%x is", helloString[i])
	} 
	fmt.Println()

	//Use utf8 package to count number of runes in the string.
	fmt.Println("Rune count is:", utf8.RuneCountInString(helloString))

	//Use range loop to handle string and decode each rune with its offset in the string
	for runeIndex, runeValue := range helloString {
		fmt.Printf("%#U starts at %d\n", runeValue, runeIndex)
	}

	//Use utf8.DecodeRuneInString function to achieve the same iteration as above
	fmt.Println("\nUsing UTF8 DecodeRuneInString Function")
	for i, w := 0, 0; i < len(helloString); i += w {
		runeValue, width := utf8.DecodeRuneInString(helloString[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width

		//Passing a rune into a function
		examineRune(runeValue)
	}

}

func examineRune(r rune) {
	//Compare a rune value to a rune literal directly
	if r == 't' {
		fmt.Println("Found the T")
	} else if r == 'こ' {
		fmt.Println("Found ko")
	}

}