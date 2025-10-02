// Replacing whole processes using Go's 'exec' function
package main

import (
	"os"
	"os/exec"
	"syscall"
)

func main() {
	//Exec dir and use 'exec.LookPath' to find the absolute path
	binary, lookupError := exec.LookPath("ls")
	if lookupError != nil {
		panic(lookupError)
	} 

	arguments := []string{"ls", "-a", "-l", "-h"}

	//Provide curerent environment variables that exec will use
	environment := os.Environ()

	//Execute syscall.Exec. If successful, process execution stops here.
	//Process is replaced by the '/bin/ls -a -l -h' process
	executionError := syscall.Exec(binary, arguments, environment)
	if executionError != nil {
		panic()

	}
}

//Using goroutines, spawning processes and exec'ing processes covers most use cases for fork which is a classic Unix function.