package main

import "fmt"

func main() {
	// pointers
	x := 0
	fmt.Println("x =",x)
	fmt.Println("Memory addr of 'x' (from main()) =",&x)

	changeXValNow(&x) // pointers are passed using &
	fmt.Println("x =",x)

	// creating new pointers using new()
	yPtr := new(int)
	changeYValNow(yPtr)
	fmt.Println("y =", *yPtr) // simply yPtr shows addr

	// structures
	rect1 := Rectangle{lextX:0, topY:50, height:10, width:10}
	rect2 := Rectangle{0, 50, 10, 10} // order allocated as in struct
	
	fmt.Println("Rectangle is",rect1.width,"unit(s) wide!")
	fmt.Println("Area of rect2 is:",rect2.rectArea())
}

// pointers in called func represented by *
func changeXValNow(x *int){
	*x = 2
	fmt.Println("Memory addr of 'x' (from changeXValNow()) =",x)
}

func changeYValNow(y *int){
	*y = 100
}

//////
type Rectangle struct {
	lextX float64
	topY float64
	height float64
	width float64
}
// assigned to the above struct
func (rect *Rectangle) rectArea() float64 {
	return rect.width * rect.height
}