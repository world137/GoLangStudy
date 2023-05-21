package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type react struct {
	width, height float64 // parameter
}
type circle struct {
	radius float64
}

func (r react) area() float64 {
	return r.width * r.height * 0.5
}
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (r react) perim() float64 {
	return 2*r.width + 2*r.height
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println("react area : ", g.area())
	fmt.Println("circle area : ", g.perim())
}
func main() {
	r := react{width: 5, height: 10}
	c := circle{radius: 20}
	measure(r)
	measure(c)
}
