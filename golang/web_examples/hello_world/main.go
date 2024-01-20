package main

import (
	"fmt"
	"net/http"
)

/*
	INTRODUCTION
	Go is a battery included programming lang, it has a webserver already built in
	The net/http package from the std library contains all functionalities about the http protocol
	This includes an HTTP client and an HTTP server
*/

func main() {
	/*
		REGISTERING A REQUEST HANDLER
		First, we create a Handler which receives all incoming HTTP connections from browsers, HTTP clients, or API reqs
		A Handler in go has the following signature:

			func(w http.ResponseWriter, r *http.Request)

		The func receives two parameters
			1. "http.ResponseWriter" is where you write your text/html response to
			2. "http.Request" contains all the information about the HTTP request (URL, header fields, etc..)
	*/

	//REGISTERING OUR HANDLER:
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Hello, you've requested: %s \n", r.URL.Path)
	})

	/*
		LISTEN FOR HTTP CONNECTIONS
		The request handler alone cannot accept any HTTP connections from the outside
		An HTTP server has to listen on a port to pass connections on to the request handler
		Port 80 is typically the default port for HTTP traffic, so that's what we're using here

	*/
	
	//LISTENING FOR CONNECTIONS
	http.ListenAndServe(":80", nil)
}