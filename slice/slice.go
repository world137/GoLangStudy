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

package main

import "fmt"

func main() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
	
	//[John Paul George Ringo]
	//[John Paul] [Paul George]
	//[John XXX] [XXX George]
	//[John XXX George Ringo]

	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s) //len=6 cap=6 [2 3 5 7 11 13]


}
func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

