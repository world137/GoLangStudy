package main

import "fmt"

func main() {
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	// while
	var text string
	for {
		fmt.Scanf("%s", &text)
		fmt.Println(text)
		if text == "exit" {
			break
		}
	}
}
