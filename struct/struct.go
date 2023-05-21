package main

import "fmt"

type car struct {
	CarID   string
	CarName string
	Color   string
}

func main() {
	car1 := car{
		CarID:   "001",
		CarName: "BMW",
		Color:   "white",
	}
	carlist := [3]car{}
	carlist[0] = car{
		CarID:   "001",
		CarName: "BMW",
		Color:   "white",
	}
	carlist[1] = car{
		CarID:   "002",
		CarName: "Benz",
		Color:   "white",
	}
	carlist[2] = car{
		CarID:   "003",
		CarName: "Volvo",
		Color:   "black",
	}
	fmt.Println("Car1 : ", car1)
	fmt.Println("Carlist : ", carlist)
}
