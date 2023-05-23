package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {

	// read the hello.txt
	content, err := ioutil.ReadFile("data.csv")

	if err != nil {
		log.Fatalf("error while reading %v", err)
	}

	// convert the byte into string
	fmt.Println(string(content))
}
