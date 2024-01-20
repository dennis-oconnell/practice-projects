package main

import (
	"fmt"
	"net/http"
)

/*
	INTRODUCTION
	This exercise is about creating a basic HTTP server in Go
	A basic HTTP server has a few jobs:

		1. PROCESSING DYNAMIC REQUESTS: process req's from users who browse the website, login into their accounts or post images

		2. SERVE STATIC ASSETS: serve Javascript, CSS, and images to browsers to create a dynamic experience for the user

		3. ACCEPT CONNECTIONS: the HTTP server must listen to a specific port to be able to accept connections from the internet

*/

func main() {
	
/*
	1. PROCESS DYNAMIC REQUESTS
	The 'net/http' package has everything we need to accept requests and handle them dynamically
	Below we will register a new handler with the 'http.HandleFunc' function
	The first parameter takes a path, and the second takes a function to execute
	For the dynamic aspect, we can read GET parameters with: 
		r.URL.Query().Get("token") 
	Or, we can read POST parameters (fields from an HTML form) with:
		r.FormValue("email")
*/
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Welcome to this website made possible by Go! HEY LOOK AT ME")
		fmt.Fprint(w, "Welcome to this website made possible by Go!")
	})

/*
	2. SERVING STATIC ASSETS
	To serve static assets like javascript, css, and images, we can use the inbuilt 'http.FileServer'
	Then we can point it to a URL path
	For a server to work, it needs to know where to serve files from
*/
	fs := http.FileServer(http.Dir("static/"))

	//Once the file server is in place, we just need to point a URL path to it
	//We do this in the same way that we did with dynamic requests
	http.Handle("/static/", http.StripPrefix("/static/", fs))


/*
	3. ACCEPT CONNECTIONS
	The last thing to finish off our basic HTTP server is to listen on a port to accept new connections from the internet
	Go has an inbuilt HTTP server
*/
	http.ListenAndServe(":80", nil)

}