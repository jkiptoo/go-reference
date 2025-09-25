//Create data that is not needed after the program exits

package main

import "fmt"

//Start by checking for errors
func checkErrors(e error) {
	if e != nil {
		panic(e)
	}
}


func main() {
	
	print := fmt.Println
	
	//Create temporary file by calling os.CreateTemp method which will open the file for reading and writing

	tempFile, err := os.Cre 
	


}