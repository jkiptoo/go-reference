// Parsing URLs in Go
package main

import (
	"fmt"
	"net"
	"net/url"
)

func main() {
	
	print := fmt.Println
	
	//Parse an example URL which contains a scheme, authentication information, host, port, path, query parameters, and query fragments. 

	urlPath := "postgres://user:pass@host.com:5432/path?k=v#f"

	//Parse URL to eliminate errors
	parseURL, err := url.Parse(urlPath)
	if err != nil {
		panic(err)
	}

	//Access Scheme
	print(parseURL.Scheme)

	//Use User method to handle all authentication information
	print(parseURL.User)
	print(parseURL.User.Username())
	userPass, _ := parseURL.User.Password()
	print(userPass)

	//Use Host method to use both hostname and the port
	print(parseURL.Path)
	
	//Use 'SplitHostPort' to extract hostname and port
	host, port, _ := net.SplitHostPort(parseURL.Host)
	print(host)
	print(port)

	//Extract the path and the fragment after the hash # sign
	print(parseURL.RawQuery)
	print(parseURL.Fragment)

	//Use RawQuery to get parameters in a string k=v format
	print(parseURL.RawQuery)
	
	//Parse query parameters into a map
	parseMap, _ := url.ParseQuery(parseURL.RawQuery)
	print(parseMap)
	print(parseMap["k"][0])
}