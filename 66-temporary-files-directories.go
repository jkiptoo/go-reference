//Create data that is not needed after the program exits

package main

import (
	"fmt"
	"os"
	"path/filepath"
)

//Start by checking for errors
func checkErrors(e error) {
	if e != nil {
		panic(e)
	}
}


func main() {
	
	print := fmt.Println
	
	//Create temporary file by calling os.CreateTemp method which will open the file for reading and writing

	tempFile, err := os.CreateTemp("", "Open sample file")
	checkErrors(err)
	
	//Display the name of the temporary file. File name starts with prefix given as second argument. The rest of the file name is auto-generated.
	print("Temporary file name is: ", tempFile.Name())
	
	//Clean up the file after usage is complete
	defer os.Remove(tempFile.Name())

	//Write some data to the file
	_, err = tempFile.Write([]byte{100, 200, 95, 85})
	checkErrors(err)

	//Create temporary directory to handle creation of many files
	temporaryDirectory, err := os.MkdirTemp("", "Temporary Directory\n")
	checkErrors(err)
	print("Temporary Directory Name is: ", temporaryDirectory)

	//Remove directory after usage is complete
	defer os.RemoveAll(temporaryDirectory)

	//Attach prefix of temporary directory to file name
	fileName := filepath.Join(temporaryDirectory, "file-one")
	err = os.WriteFile(fileName, []byte{7,8}, 0666)
	checkErrors(err)

}