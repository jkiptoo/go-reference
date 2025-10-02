// Using net/http package to create HTTP Server
package main

import (
	"fmt"
	"net/http"
)

//Implement handler by applying 'http.HandlerFunc' on functions
func basicHandler(responseWriter http.ResponseWriter, request *http.Request) {
	
	//Fill in the HTTP response.
	//'http.ResponseWriter and 'http.HandlerFunc' are passed as arguments to functions that act as handlers.
	
	fmt.Fprintf(responseWriter, "Server is alive\n")

	}

//Read all HTTP request headers and output them in the response body
func headerHandler(responseWriter http.ResponseWriter, request *http.Request) {
	for name, headerHandler := range request.Header {
		for _, headerOutput := range headerHandler {
			fmt.Fprintf(responseWriter, "%v: %v\n", name, headerOutput)
		}
	}
}


func main() {
	//Setup default router taking a function as a default
	// Register handlers on server routes using http.HandleFunc method
	http.HandleFunc("/server-message", basicHandler)
	http.HandleFunc("/server-header", headerHandler)

	//Call 'ListenAndServe' with port and handler. 
	//Use nil to tell router to use the default router that has been setup.
	http.ListenAndServe(":8090", nil)


}

//-----NOTES-----//
//Access headerHandler route using 'curl localhost:8090/headerHandler'