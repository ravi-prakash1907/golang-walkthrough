package main

import "fmt"

var print = fmt.Printf
var input = fmt.Scanf

func main() {
	var num int

	print("Enter number of fibonacci terms you want: ")
	input("%d", &num)

	print("\n------------\n\n")
	fibSeries(num)
	print("\n")

}

func nthFibNumber(num int) int {
	if num == 0 || num == 1 {
		return num
	} else {
		return nthFibNumber(num-1) + nthFibNumber(num-2)
	}
}

func fibSeries(length int) {
	var nth_term int

	if length > 0 {
		for n := 0; n < length; n++ {
			nth_term = nthFibNumber(n)
			print("Term %d = %d\n", n+1, nth_term)
		}
	} else {
		print("Invalid input size for series! Try later!!\n")
	}
}