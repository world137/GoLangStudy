package main

import "fmt"

func main() {
	var name = make(map[string]float64) // key value
	fmt.Println(name)

	//add
	name["a"] = 4000.00
	name["b"] = 3000.00
	name["c"] = 6500.00
	fmt.Println(name)

	//del
	delete(name, "c")
	fmt.Println(name)

	//update
	name["b"] = 2000.40
	fmt.Println(name)

	value := name["a"]
	fmt.Println(value)

	newname := map[string]string{"a": "400", "b": "200", "c": "899"}
	fmt.Println(newname)
}
