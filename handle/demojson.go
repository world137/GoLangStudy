package main

import (
	"encoding/json"
	"fmt"
)

type employee struct {
	// ชื่อเป็นตัวพิมพ์ใหญ่ => json.Marshal อ่านเฉพาะ key ที่เป็นตัวพิมพ์ใหญ่
	ID   int
	Name string
	Tel  string
	Mail string
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
func main() {
	data, err := json.Marshal(&employee{100, "world", "0213918999", "world@gmail.com"})
	check(err)
	fmt.Println(string(data))

}
