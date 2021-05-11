package main

// importing multiple libraries
import (
	"fmt"
	"os"
	"log"
	"io/ioutil"
	"strconv"
	"time"
)

// global variables
var pizzaNum = 0
var pizzaName = ""

func main() {

	fileName := "assets/lol.txt"
	// creating a file
	fileMaker("Here we go!!", fileName)

	// reading a file
	fileViewer(fileName)

	// conversion
	conversions()

	/// goroutine
	// pausing the exec : working with 'time' lib
	for i := 0; i < 10; i++{
		// a goroutine : it's executed by go 
		go count(i)
	}

	time.Sleep(time.Millisecond * 11000)
	fmt.Println("READ MORE ABOUT GOROUTINE!\n")

	// passing data b/w goroutine
	strChan := make(chan string)

	for i:=0; i < 3; i++ {
		go makeDough(strChan)
		go addSauce(strChan)
		go addToppings(strChan)

		time.Sleep(time.Millisecond * 5000)
	}
}

////////////////////
func fileMaker(gtoStr, fName string) {
	file, err := os.Create(fName)

	if err != nil {
		log.Fatal(err) // in logs
	}

	file.WriteString(gtoStr)
	file.Close()
	fmt.Println("File Created successfully!")
}

func fileViewer(fName string) {
	stream, err := ioutil.ReadFile(fName)
	if err != nil {
		log.Fatal(err) // in logs
	}

	readStr := string(stream)
	fmt.Println("\nRead from file: \n",readStr)
}
////////////////////

func conversions() {
	// to b converted
	randInt := 465
	randFloat := 323.5
	randStr1 := "100"
	randStr2 := "789.23"

	// conversion
	fmt.Println(float64(randInt))
	fmt.Println(int(randFloat))

	newInt, _ := strconv.ParseInt(randStr1, 0, 64) // 0, 64-bit
	fmt.Println(newInt)
	newFloat, _ := strconv.ParseFloat(randStr2, 64) // 64-bit
	fmt.Println(newFloat)
}

func count (id int) {
	for i := 0; i < 10; i++ {
		fmt.Println(id,":", i)

		// pauing for a sec
		time.Sleep(time.Millisecond * 1000)
	}
}

////////////////////
func makeDough(strChan chan string) {
	pizzaNum++

	pizzaName = "Pizza #" + strconv.Itoa(pizzaNum)
	fmt.Println("\nMake dough for", pizzaName, "and send for sauce")

	// passing to next goroutine
	strChan <- pizzaName
	time.Sleep(time.Millisecond * 10)
}

func addSauce(strChan chan string) {
	pizza := <- strChan

	fmt.Println("Add sauce to", pizza, "and send for toppings")

	// passing to next goroutine
	strChan <- pizzaName
	time.Sleep(time.Millisecond * 10)
}

func addToppings(strChan chan string) {
	pizza := <- strChan

	fmt.Println("Add toppings to", pizza, "and ship")

	time.Sleep(time.Millisecond * 10)
}
////////////////////