package main

// importing multiple libraries
import (
	"fmt"
	"net/http"
)

// global variables
var pizzaNum = 0
var pizzaName = ""

func main() {
	// HTTP server
	fmt.Printf("Server is up")
	fmt.Printf("Access at 'localhost:8080' and 'localhost:8080/earth'\n")
	fmt.Println("Stop the server using 'Ctrl/Cmd + C'")

	// request will go into '/' dir and 'handler wil  handel it
	http.HandleFunc("/", handler) 
	// request will go into '/earth' dir and 'handler2 wil  handel it
	http.HandleFunc("/earth", handler2)

	// listne on specified port
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World\n")
}

func handler2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Earth\n")
}