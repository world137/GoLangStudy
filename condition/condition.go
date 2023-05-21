package main

import "fmt"

func cal(i int) {
	if i >= 80 {
		fmt.Println("A")
	} else if i >= 70 {
		fmt.Println("B")
	} else if i >= 60 {
		fmt.Println("C")
	} else if i >= 50 {
		fmt.Println("D")
	} else {
		fmt.Println("F")
	}
}

func main() {

	var i int

	fmt.Print("Type a number: ")
	fmt.Scan(&i)
	cal(i)
}
