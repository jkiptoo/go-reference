//Reading files

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

//Use helper to streamline error checks
func checkErrors(e error) {
	if e != nil {
		panic(e)
	}
}


func main() {
	
	print := fmt.Println
	//Transfer file's contents into memory
	data, err := os.ReadFile("C:/wamp64/www/Go/exercises/basic/go-reference/tmp/data.txt")
	checkErrors(err)
	print(string(data))

	//Assert control over sections of file to read
	//Open file to get os.File value
	openFile, err := os.Open("C:/wamp64/www/Go/exercises/basic/go-reference/tmp/data.txt")
	checkErrors(err) 

	//Read upto 5 bytes from the start of a file
	readBytes := make([]byte, 5)
	checkReadBytes, err := openFile.Read(readBytes)
	checkErrors(err)
	print("%d bytes: %s\n", checkReadBytes, string(readBytes[:checkReadBytes]))

	//Move to a known location in the file and read from there
	knownLocation, err := openFile.Seek(6, io.SeekStart)
	checkErrors(err)
	secondReadBytes := make([]byte, 2)
	secondKnownLocation, err := openFile.Read(secondReadBytes)
	checkErrors(err)
	print("%d bytes @ %d: ", secondKnownLocation, knownLocation)
	print("%v\n", string(secondReadBytes[:secondKnownLocation]))

	//Seek to the current cursor position
	_, err = openFile.Seek(2, io.SeekCurrent)
	checkErrors(err)

	
	//See to the end of the file
	_, err = openFile.Seek(-4, io.SeekEnd)
	checkErrors(err)

	//Using 'ReadAtLeast' to implement more robust file reads
	thirdKnownLocation, err := openFile.Seek(6, io.SeekStart)
	checkErrors(err)
	thirdReadBytes := make([]byte, 2)
	readNow, err := io.ReadAtLeast(openFile, thirdReadBytes, 2)
	checkErrors(err)
	fmt.Printf("%d bytes @ %d: %s\n", readNow, thirdKnownLocation, string(thirdReadBytes))

	//Rewind using 'Seek(0, io.SeekStart)'
	_, err = openFile.Seek(0, io.SeekStart)
	checkErrors(err)

	//Use bufio to implement buffere reader. bufio provides additional reading methods
	readBuffer := bufio.NewReader(openFile)
	fourthReadBytes, err := readBuffer.Peek(5)
	checkErrors(err)
	fmt.Printf("5 bytes: %s\n", string(fourthReadBytes))

	//Close file after finishing reading
	openFile.Close()
	
}