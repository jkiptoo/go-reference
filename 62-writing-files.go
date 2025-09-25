// Writing files in go
package main

import (
	"bufio"
	"fmt"
	"os"
)

//Check errors
func errorCheck(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	print := fmt.Printf
	//Dump a string into a file
	stringDump := []byte("where is\ngolang\n")
	err := os.WriteFile("C:/wamp64/www/Go/exercises/basic/go-reference/tmp/writing.txt", stringDump, 0644)
	errorCheck(err)

	//Open file for reading
	openFile, err := os.Create("C:/wamp64/www/Go/exercises/basic/go-reference/tmp/write.txt")
	errorCheck(err)

	//Defer close immediately after opening a file
	defer openFile.Close()

	//Write a byte slice
	stringByte := []byte{15, 20, 25, 30, 65}
	writeBytes, err := openFile.Write(stringByte)
	errorCheck(err)
	print("Already written %d bytes\n", writeBytes)

	//Use 'WriteString' instead
	secondWriteBytes, err := openFile.WriteString("Writing in progress\n")
	errorCheck(err)
	print("Written %d bytes\n", secondWriteBytes)

	//Flush writes into a stable storage using Sync
	openFile.Sync()

	//Use bufio writers 
	writeBuffer := bufio.NewWriter(openFile)
	thirdWriteBytes, err := writeBuffer.WriteString("Buffer writes\n")
	errorCheck(err)
	print("Buffer has written %d bytes\n", thirdWriteBytes)
	writeBuffer.Flush()

}