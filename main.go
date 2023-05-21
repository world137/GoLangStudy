package main

import "fmt"

var numberInt int = 1000
var smg string = "hello"

func main() {
	floatNum := 20.00
	fmt.Println(numberInt)
	fmt.Println(float64(numberInt) + floatNum)
	fmt.Println(smg + "world")
	fmt.Println("money", floatNum)
}
