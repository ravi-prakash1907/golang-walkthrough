// this pkg
package main

import "fmt"

func main(){
	// printing regilar output
	fmt.Println("Hello World!")

	// declearing primary variables and constants
	basicDeclearations()

	// using common operators in go
	operators()
	
	// formatted outputs
	formattings()

	//Logic
	logicals()

	//Loop
	loops()

	//Relational 
	relationals()

	// Conditional stmsts.
	conditionals()

	// Derived datatypes
	derivedsVar()

}

func basicDeclearations(){
	// variables --> various types
	// the type can't be changed now
	var age int = 40
	
	var favNum float64 = 1.7294471;  
	// ; is not necessary at end, but puttin it won't cause a trouble

	randNum := 1  // automatically int 

	// randNum = "Hi" // this is error
	fmt.Println("",age,"\n",favNum,"\n",randNum)

	// constants
	const pi float64 = 3.14159265

	// declearing multiple variables
	var (
		varA = 2
		varB = 3
	)

	var myName string = "Ravi Prakash"
	fmt.Println("My name is "+myName,"\nlength of name is:",len(myName))
	fmt.Println("varA x varB = ",varA*varB)
}

func operators(){
	var numOne = 1.0000
	var num99 = .9999
	fmt.Println(numOne - num99) // not very accurate

	// aruthmatic
	num1 := 24
	num2 := 6
	fmt.Println(num1 + num2)
	fmt.Println(num1 - num2)
	fmt.Println(num1 * num2)
	fmt.Println(num1 / num2)
	fmt.Println(num1 % num2)
}

func formattings(){
	// prior work
	const pi float64 = 3.14159265
	var (
		varA = 2
		varB = 3
	)

	// bools
	var flag bool = true 
	fmt.Printf("Pi = %f \n", pi)
	fmt.Printf("Pi upto 3 decimals = %.3f \n", pi)

	// get datatype
	fmt.Printf("%T \n",varA)
	fmt.Printf("%T \n",varB)
	fmt.Printf("%T \n",flag)

	// getting boolean res, usage of %t
	fmt.Printf("%t \n",flag)

	///////////////////////
	fmt.Printf("%d \n", 100)  // decimal of 100
	fmt.Printf("%b \n", 100)  // bin of 100
	fmt.Printf("%c \n", 44)   // char for ascii val 44
	fmt.Printf("%x \n", 17)   // hex of 17
	fmt.Printf("%o \n", 17)   // oct of 17
	fmt.Printf("%e \n", pi)   // sci. notation of pi
}

func logicals(){
	//and
	fmt.Println("true && false = ", true && false)
	//or
	fmt.Println("true || false = ", true || false)
	//not
	fmt.Println("!true = ", !true)
}

func loops(){
	// for loop
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++  // similarly i-- 
	}
}

func relationals(){
	/*
	following are the relation operators in go
		==
		< 
		>
		!=
		<=
		>=
	*/
	for j := 0; j < 5; j++ {
		fmt.Println(j)
	}
}

func conditionals(){
	// if-else
	age := 18

	if age >= 16 {
		fmt.Println("You can drive!")
	} else if age >= 18 {
		fmt.Println("You can vote!") // won't execute
	} else {
		fmt.Println("You can have fun!")
	}

	// switch
	age = 5

	switch age {
	case 16 : fmt.Println("You can drive!")
	case 18 : fmt.Println("You can vote!")
	default : fmt.Println("You can have fun!")
	}
}

func derivedsVar(){
	// array - method 1
	var favNum2[5] float64

	favNum2[0] = 267
	favNum2[1] = 789287
	favNum2[2] = 27
	favNum2[3] = 53.5
	favNum2[4] = 2.34

	fmt.Println(favNum2[3])

	// array - method 2
	favNum3 := [5]float64 {1,2,3,4,5}
	// i is index, val is data
	fmt.Println("With index:")
	for i, val := range favNum3 {
		fmt.Println(val, i)
	}

	// get rid of index 
	fmt.Println("Without index:")
	for _, val := range favNum3 {
		fmt.Println(val)
	}


	// slices:  may have blank space
	// i.e. size not decleared
	numSlice := []int {2,1,4,6,7}
	numSlice2 := numSlice[3:5]  // [:index], [index:]

	fmt.Println("numSlice2[0] =", numSlice2[0])

	// declearing slice using make(type, defaultZeroTOHowManyInitialNum, maxSize)
	numSlice3 := make([]int, 5, 10)

	// copying
	copy(numSlice3, numSlice2)
	fmt.Println(numSlice3[0])

	// append(slice, index among blank spaces, val)
	numSlice3 = append(numSlice3, 0, -1) // here 0 means 6th place.. guess why?
	fmt.Println(numSlice3[6])
}