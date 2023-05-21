package main

import "fmt"

func test() {
	fmt.Println("hello world")
}
func cal(a int, b int) {
	sum := a + b
	fmt.Println(sum)
}
func returncal(a int, b int) float64 { // (val1,val2) return type
	return float64(a + b)
}

func main() {
	test()
	cal(10, 20)
	smg := returncal(20, 20)
	fmt.Println(smg)
}
