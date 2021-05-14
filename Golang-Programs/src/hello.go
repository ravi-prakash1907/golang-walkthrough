package main

import "fmt"

func main() {
    var fname string

    fmt.Printf("Please enter your first name: ")
    fmt.Scanf("%s", &fname)

    if len(fname) > 0 {
        fmt.Printf("Hello "+fname+"!\n")
    } else {
        fmt.Printf("Hello World!\n")
    }    
}