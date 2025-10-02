//Using 'log' package for free-form output.
//Using 'log/slog' for structured output.

package main

import (
	"bytes"
	"fmt"
	"log"
	"log/slog"
	"os"
)

func main() {
	//Invoke functions that have already been pre-configured for reasonable logging output. Output is sent to os.Stderr
	//Fatal* or Panic* methods exit the program after logging.
	log.Println("This is the Standard Logger.")
	
	//Configure loggers to set their output format. e.g. change output Ldate and Ltime to not possess time and microsecond accuracy.
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.Println("With Microseconds.")

	//Emit file name and line then call log function.
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("With file or line.") 

	//Create a custom logger to pass aroungd the program by using a prefix to make its output appear different from othe loggers.
	logOutput := log.New(os.Stdout, "My:", log.LstdFlags)
	logOutput.Println("This is my log...")

	//Using SetPrefix method to set a prefix on existing loggers
	logOutput.SetPrefix("For my eyes only:")
	logOutput.Println("From System Log")

	//Use io.Writer to provide custome targets for logger outputs
	var customBuffer bytes.Buffer
	bufferLog := log.New(&customBuffer, "Custom Buffer:", log.LstdFlags)

	//Write output from log into buffer
	bufferLog.Println("Greetings from the buffer zone.")

	//Write output above into standard output
	fmt.Print("From Buffer Log:", customBuffer.String())

	//Use slog package to provide structured log output e.g. logging in JSON format
	jsonHandler := slog.NewJSONHandler(os.Stderr, nil)
	slogOutput := slog.New(jsonHandler)
	slogOutput.Info("Here is the JSON handler.")

	//Display 'key=value' pairs using the slog package
	slogOutput.Info("Here is the slog package output", "key", "value", "age", 54)

}