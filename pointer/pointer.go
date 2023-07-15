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



	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	fmt.Println(p) // address
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}
