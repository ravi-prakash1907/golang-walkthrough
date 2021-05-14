package main

import "fmt"

var print = fmt.Println
var scan = fmt.Scanf

func main() {
	var num int

	print("Enter a number to get its factorial: ")
	scan("%d", &num)

	result := fact(num)
	print(num,"factorial =",result)
}

func fact (n int) int {
	if n <= 1 {
		return 1
	} else {
		return n * fact(n-1)
	}
}