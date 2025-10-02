// Generating other processes within  the Go program
// Spawn processes so as to make external processes accessible to a running Go process
package main

import (
	"errors"
	"fmt"
	"io"
	"os/exec"
)

func main() {
	
	print := fmt.Println
	//Create a command that takes no arguments nor input and outputs to stdout
	//'exec.Command' helper creates an object to demonstrate this external process
	dateCommand := exec.Command("date")
	
	//Run the command, wait for it to finish then gather its standard output.
	dateOutput, err := dateCommand.Output()
	if err != nil {
		panic(err)
	}
	//If there are no errors, 'dateOutput' will hold bytes with the date information
	print("> date")
	print(string(dateOutput))

	//If the execution of the program is encountered, 'Output' and other methods will return *exec.Error
	_, err = exec.Command("date", "-x").Output()
	if err != nil {
		var executionError *exec.Error
		
		//If the command ran but exited with a zero return code, '*exec.ExitError' will be returned
		var exitError *exec.ExitError
		
		switch {
		case errors.As(err, &executionError):
			print("Failed to execute the command:", err)
		case errors.As(err, &exitError):
			exitCode := exitError.ExitCode()
			print("The return code produced upon command exit is =", exitCode)
		default:
			panic(err)
		}
	}
	//Pipe data to an external process in its 'stdin' and collect its output from its 'stdout'

	grepCommand := exec.Command("Grep says", "server is loading...")

	//Directly get input and output pipes (Add error handling)
	grepInput, _ := grepCommand.StdinPipe()
	grepOutput, _ := grepCommand.StdoutPipe()

	//Start the process
	grepCommand.Start()

	//Write input into the process 
	grepInput.Write([]byte("Grep is starting\nGrep is saying goodbye. Now shutting down."))
	grepInput.Close()

	//Read the resulting output (Add error handling)
	grepBytes , _ := io.ReadAll(grepOutput)
	
	//Wait for the process to exit
	grepCommand.Wait()

	//Now collecting 'StdoutPipe' results. 
	//'StderrPipe' results can be collected in the same way
	print("> Grep says I am now starting...")
	print(string(grepBytes))

	//SPAWNING RULE: Supply delineated command and argument array
	//Spawn a full command with a string using Bash's -c option.
	spawnCommand := exec.Command("Bash Shell", "-c", "ls -a -l -h")
	spawningOutput, err := spawnCommand.Output()
	if err != nil {
		panic(err)
	}
	print("> ls -a -l -h")
	print(string(spawningOutput))

}

//----------------NOTES---------------//
//Spawned programs return output same way if they were run from the ommand line.
//Program exits with an error message and non-zero return code because 'date' does not have a -x flag.