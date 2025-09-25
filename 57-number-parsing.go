// Parsing numbers from strings in Go using strconv package
package main

import (
	"fmt"
	"strconv"
)

func main() {
	
	print := fmt.Println
	//64 indicates the number of bits of precision of the parse
	parseFloats, _ := strconv.ParseFloat("3.148", 64)
	print(parseFloats)

	//0 indicates base from the string. Result should fit within 64 bits
	parseIntegers, _ := strconv.ParseInt("314", 0, 64)
	print(parseIntegers)
	
	//Use 'ParseInt' to parse hex numbers
	parseHexNumbers, _ := strconv.ParseInt("0x5d89", 0, 64)
	print(parseHexNumbers)

	//Parse unsigned integers
	parseUnsignedIntegers, _ := strconv.ParseUint("987", 0, 64)
	print(parseUnsignedIntegers)
	
	//Parse base-10 integer using 'Atoi' function
	parseBaseTen, _ := strconv.Atoi("145")
	print(parseBaseTen)
	
	//Return an error if a bad input
	_, e := strconv.Atoi("What")
	print(e)
	
}