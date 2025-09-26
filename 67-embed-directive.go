// Using embed directive to create files and folders during build runtime in the Go binary.
package main

import "embed"

//Embed contents of the file into a string variable
//go:embed folder/file.txt
var fileString string

//Embed file's contents into a []byte
//go:embed folder/file.txt
var fileByte []byte

//Embed multiple files and folders using wildcards. This implements a virtual file system.
//go:embed folder/file.txt
//go:embed folder/*.hash
var folder embed.FS

func main() {
	//Print out contents of a single file
	print(fileString)
	print(string(fileByte))

	//Retrieve embedded folder
	contentOne, _ := folder.ReadFile("folder/file-one.hash")
	print(string(contentOne))

	contentTwo, _ := folder.ReadFile("folder/file-two.hash")
	print(string(contentTwo))	
}