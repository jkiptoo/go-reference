// Using net/http package to issue HTTP requests
package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {

	print := fmt.Println
	//Issue a HTTP GET request to a server
	response, err := http.Get("https://wiredtendon.com")
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	//Print HTTP Response Status
	print("Response Status:", response.Status, response.StatusCode, response.TLS)

	//Print the first 7 lines of the response body
	responseScanner := bufio.NewScanner(response.Body)
	for i := 0; responseScanner.Scan() && i < 7; i++ {
		print(responseScanner.Text())
	}

	if err := responseScanner.Err(); 
	err != nil {
		panic(err)
	}

}