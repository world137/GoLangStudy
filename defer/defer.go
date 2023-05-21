package main

import "fmt"

func add(value1, value2 float64) {
	sum := value1 + value2
	fmt.Println("value : ", sum)
}

func main() {
	defer add(20, 20) // ทำทีหลัง
	add(10, 20)

	// lifo last in first out
	defer add(1, 2)
	defer add(3, 4)
	defer add(5, 6)
}
