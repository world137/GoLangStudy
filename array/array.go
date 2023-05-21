package main

import (
	"fmt"
	"strconv"
)

var name [4]string

func main() {
	for i := 0; i < len(name); i++ {
		name[i] = strconv.Itoa(i)
	}
	fmt.Println(name)
}
