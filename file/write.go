package main

import "os"

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	data := []byte("Hello world!")
	err := os.WriteFile("test.txt", data, 0644)
	check(err)
	f, err := os.Create("name")
	check(err)
	defer f.Close()
	data2 := []byte("world\n chayapol")
	os.WriteFile("name.txt", data2, 0644)
}
