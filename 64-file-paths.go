//Parse and construct file paths to be compatible across operating systems

package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {

	print := fmt.Println
	//Use 'Join' to construct file paths in a portable manner. Many arguments can be used
	portableFilePath := filepath.Join("directoryOne", "directoryTwo", "filename")
	print("File Path:", portableFilePath)

	//Use Join instead of concatenating. 
	//Join helps normalise paths by removing separators and directory changes
	print(filepath.Join("directoryOne//", "filename"))
	print(filepath.Join("directoryOne/../directoryTwo", "filename"))

	//Use Dir and Base to split path to directory
	//Split can do the two calls in one
	print("Dir(portableFilePath):", filepath.Dir(portableFilePath))
	print("Base(portableFilePath):", filepath.Base(portableFilePath))

	//Check if path is absolute
	print(filepath.IsAbs("dir/file"))
	print(filepath.IsAbs("/dir/file"))
	fileName := "config.json"
	
	//Split extensions out of file names
	splitExtension := filepath.Ext(fileName)
	print(splitExtension)

	//Use strings.TrimSuffix to find filename without the extension
	print(strings.TrimSuffix(fileName, splitExtension))

	//Find relative path between base and target
	//Return error if target cannot be relative to the base
	relativePath, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}

	print(relativePath)




}