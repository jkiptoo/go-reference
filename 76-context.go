// Implement request-scoped values across API boundaries and goroutines.
// Controlling cancellation using context.Context method
package main

import (
	"fmt"
	"net/http"
	"time"
)

//Create context.Context for each request using the Context() method
func serverMessage(writer http.ResponseWriter, request *http.Request) {
	print := fmt.Println
	requestContext := request.Context()
	print("Server says: Server Message Handler has Started...")
	defer print("Server says: Server Message Handler has Stopped!")

	//To simulate work that the server is doing, wait a few seconds before sending a response to client
	select {
	case <-time.After(12 * time.Second):
		fmt.Fprintf(writer, "Here is the server message after a few seconds...")
	
	//Check context's 'Done()' channel for a signal to cancel the work and return as soon as possible.
	case <-requestContext.Done():

		//Use context's 'Err()' method to return an error that explains why the 'Done()' channel was closed
		err := requestContext.Err()
		print("Server encountered Error and says:", err)
		internalServerError := http.StatusInternalServerError
		http.Error(writer, err.Error(), internalServerError)
	}
}

func main() {
	//Register handle on the 'server-status' route then start serving
	http.HandleFunc("/server-status", serverMessage)
	http.ListenAndServe(":8090", nil)
	
}

//----------NOTES----------//
//Simulate a client request to '/server-status' 
// Use $ curl 'localhost:8090/server-status'