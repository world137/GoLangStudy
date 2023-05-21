package main

import (
	"fmt"
	"time"
)

// parallel work

func print(text string) {
	for i := 0; i < 100; i++ {
		fmt.Println(text, " : ", i)
	}
}

func main() {
	// print("hello world")
	// print("my name is chayapol")

	go print("hello world")
	go print("my name is chayapol")
	time.Sleep(5 * time.Second)
}
