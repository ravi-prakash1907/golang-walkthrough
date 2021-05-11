package main

import "fmt"

func main() {

	// Maps :  as dictionary
	maps()

	// func in func :   Clouser
	num3 := 3
	myclouser(num3)

	// Recursion
	num4 := 5
	gotFact := factorial(num4)
	fmt.Println("5! =",gotFact)

	// defer
	fmt.Println("\n'defer' executes a function after ENCLOSING function ( here main() ) finishes its execution\n")
	defer printTwo()
	printOne()

	// catching errors : recovering
	fmt.Println("recover() allows to continue execution even if an error occures in b/w")
	fmt.Println(safeDevision(6,0))
	fmt.Println(safeDevision(6,3))

	// panic : for error handling
	demoPanic()

}

func maps() {
	// maps are the key : value pairs
	presAge := make(map[string] int)

	presAge["first"] = 42
	fmt.Println(presAge["first"]," | ",len(presAge))

	presAge["John F. Kennedy"] = 43
	fmt.Println(presAge["John F. Kennedy"]," | ",len(presAge))

	// deleting items
	delete(presAge, "John F. Kennedy")
	fmt.Println(len(presAge))

	// levelUp function calls
	listNums := []float64 {1,2,3,4,5}
	fmt.Println("Sum :", allAll(listNums))

	// return multiples
	num1, num2 := next2Vals(5)
	fmt.Println(num1,num2)

	// Variadic function -> undefined parameters
	fmt.Println(subtractThem(1,2,3,4,6))
}

///////////////
func allAll(numbers []float64) float64 {
	sum := 0.0 // must be float, else wont be added

	for _, item := range numbers {
		sum += item //similarly -= ...
	}

	return sum
}

func next2Vals(x int) (int, int){
	return x+1, x+2
}

func subtractThem(args ...int) int {
	finalVal := 0

	for _, val := range args {
		finalVal -= val
	}

	return finalVal
}
///////////////

func myclouser(num int) {
	
	doubleNum := func() int {
		num *= 2
		return num	
	}

	fmt.Println(doubleNum())
	fmt.Println(doubleNum())
}

func factorial(num int) int {
	if num == 0 {
		return 1
	} else {
		return num*factorial(num - 1)
	}
}

///////////
func printOne(){fmt.Println(1)}
func printTwo(){fmt.Println(2)}
///////////

func  safeDevision(num1, num2 int) int {
	defer func() {
		// prints error if any
		// just call recover() to avoid error prompt
		fmt.Println(recover()) 
	}()

	solution := num1 / num2
	return solution
}

func demoPanic(){
	defer func(){
		// if Println is removed, nothing is printed
		// else will print PANIC
		fmt.Println(recover())
	}()

	panic("PANIC")
}