package main

import "fmt"

func main() {
	// input := 3
	// switch input {
	// case 1:
	// 	fmt.Println("1")
	// case 2:
	// 	fmt.Println("2")
	// case 3:
	// 	fmt.Println("3")
	// case 4:
	// 	fmt.Println("4")
	// default: // ถ้าไม่ตรงเงื่อนไขไหนเลย
	// 	fmt.Println("none")
	// }
	var color string
	fmt.Println("color:")
	fmt.Scanf("%s", &color)
	switch color {
	case "blue":
		fmt.Println("#0000FF")
	case "green":
		fmt.Println("#008000")
	case "pink":
		fmt.Println("#FFC0CB")
	case "yellow":
		fmt.Println("#FFFF00")
	default:
		fmt.Println("error")
	}
}
