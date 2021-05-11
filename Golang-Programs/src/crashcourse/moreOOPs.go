package main

// importing multiple libraries
import (
	"fmt"
	"math"
	"strings"
	"sort"
)

func main() {
	// polymorphism using interfaces : works kinda template
	rect := Rectangle{20,50}
	circ := Circle{10}

	fmt.Println("Area of rectangle:",getArea(rect),"sq units")
	fmt.Printf("Area of circle: %.3f sq units\n",getArea(circ))

	// working with strings
	sampleStr := "Hello World"

	fmt.Println(strings.Contains(sampleStr, "lo")) // sub str. check
	fmt.Println(strings.Index(sampleStr, "lo")) // sub str. index if exists
	fmt.Println(strings.Count(sampleStr, "l")) // sub str. how many times if exists
	fmt.Println(strings.Replace(sampleStr, "l", "x", 3)) // (in, to, by, first?occurances)

	// comma seperated str
	csvStr := "1,2,3,4,5"
	fmt.Println(strings.Split(csvStr, ",")) // returns kinda list

	letterList := []string{"c", "a", "b"}
	sort.Strings(letterList) // sorting string array
	fmt.Println(letterList)

	numList := strings.Join([]string {"3","2","1"}, ", ") // Join(list, seperator)
	fmt.Println(numList)
}

type Shape interface {
	area() float64
}

type Rectangle struct {
	height float64
	width float64
}
func (r Rectangle) area() float64 {
	return r.height * r.width
}

type Circle struct {
	radius float64
}
func (c Circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

// gives area() for respective shape
func getArea(shape Shape) float64{
	return shape.area()
}