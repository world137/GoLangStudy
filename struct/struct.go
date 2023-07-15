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
type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit x:0 => x =0
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)

func main() {
	fmt.Println(v1, p, v2, v3)
	fmt.Println(v2.X)
	fmt.Println(v2.Y)
}
