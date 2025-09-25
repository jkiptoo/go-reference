// Working with directories
package main

import (
	"io/fs"
	"os"
	"path/filepath"
)

//Check for errors
func checkErrors(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	//Create new subdirectory in the current working directory
	err := os.Mkdir("subdir", 0755)
	checkErrors(err)
	
	//Defer removal when creating temporary directory (rm -rf)
	defer os.RemoveAll("subdir")

	//Use helper function to create empty file
	createEmptyFile := func(name string) {
		createdFile := []byte("")
		checkErrors(os.WriteFile(name, createdFile, 0644))
	}
	createEmptyFile("subdir/fileOne")

	//Create a hierarchy of directories including MkdirAll
	err = os.MkdirAll("subdir/parent/child", 0755)
	checkErrors(err)

	createEmptyFile("subdir/parent/fileTwo")
	createEmptyFile("subdir/parent/fileThree")
	createEmptyFile("subdir/parent/child/fileFour")

	//List directory contents returning a slice of os.DirEntry objects
	directoryContent, err := os.ReadDir("subdir/parent")
	checkErrors(err)

	print("Listing subdir/parent\n")
	for _, entry := range directoryContent {
		print(" ", entry.Name(), entry.IsDir())
	}

	//Change current working directory using Chdir
	err = os.Chdir("subdir/parent/child")
	checkErrors(err)

	//View contents of subdir/parent/child when listing current directory
	directoryContent, err = os.ReadDir(".")
	checkErrors(err)

	print("Listing subdir/parent/child\n")
	for _, entry := range directoryContent {
		print(" \n", entry.Name(), entry.IsDir())
	}

	//Change directory back to original starting point
	err = os.Chdir("../../..")
	checkErrors(err)

	//Transverse directory recursively plus subdirectories
	print("Transversing subdir\n")
	err = filepath.WalkDir("subdir", transverse)
	
}

func transverse(path string, createdFile fs.DirEntry, err error) error {
	if err != nil {
		return err
	}
	print(" \n", path, createdFile.IsDir())
	return nil
}