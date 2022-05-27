package main

import (
	"fmt"
	"net/http"
)

// The below code creates a "handler", I need to figure out what that is.
// In the below handler() function, "w" is the response writer, "r" is the request. (Thanks github copilot)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

// Below, the http.HandleFunc() function takes two parameters:
// 1. The path to the handler function. In this case, it is "/" (GHC).
//    In the function description, that parameter is listed as 'pattern'.
//    I don't know what a pattern is for now. NH
// 2. The handler function itself (GHC).

// http.ListenAndServe() seems to be some sort of function that communicated with server.
// It takes two parameters:
// 1. The port number to listen on. (GHC)
//    In the function description, that parameter is listed as 'addr'. (GHC)
// 2. The handler function. (GHC)
//    Description says this parameter is usually nil. NH
//	  This still confuses me. Both http.HandleFunc() and http.ListenAndServe() take a handler as the second parameter . NH
//    So why dont they accept the same handler?

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
