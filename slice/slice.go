package main

import (
	"fmt"
)

func main() {
	var name []string
	name = []string{"a", "b"}
	fmt.Println(name)
	name = append(name, "c")
	fmt.Println(name)
	var name2 []string
	name2 = []string{"world", "hello"}
	name2 = append(name2, name...)
	fmt.Println(name2)
	newname := name2[1:3]
	fmt.Println(newname)

}
