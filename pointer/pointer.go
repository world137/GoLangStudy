package main

import "fmt"

func setZero(val int) {
	val = 0
}
func setZeroPointer(val *int) {
	*val = 0
}

func main() {
	i := 1
	fmt.Println(i)
	setZero(i)
	fmt.Println(i)
	setZeroPointer(&i) // ==> & กำหนด address pointer
	fmt.Println(i)
}
