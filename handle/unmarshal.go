package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type employee struct {
	ID   int
	Name string
	Tel  string
	Mail string
}

func main() {
	e := employee{}
	err := json.Unmarshal([]byte(`{"ID":101,"Name":"world","Tel":"58293457","Mail":"world@gmail.com"}`), &e)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(e)
	fmt.Println(e.Name)
}
