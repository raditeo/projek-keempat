package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
	perimeter() float64
}

type rectangle struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (r rectangle) area() float64 {
	return r.height * r.width
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (r rectangle) perimeter() float64 {
	return 2 * (r.height * r.width)
}

func (c circle) volume() float64 {
	return 4 / 3 * math.Pi * math.Pow(c.radius, 3)
}

func print(t string, s shape) {
	fmt.Printf("%s area %v\n", t, s.area())
	fmt.Printf("%s perimeter %v\n", t, s.perimeter())
}

func main() {
	var c1 shape = circle{radius: 14}
	var r1 shape = rectangle{width: 5, height: 7}

	// fmt.Printf("type of c1: %T\n", c1)
	// fmt.Printf("type of r1: %T\n", r1)

	// fmt.Println("circle area:", c1.area())
	// fmt.Println("circle perimeter:", c1.perimeter())

	// fmt.Println("rectangle area:", r1.area())
	// fmt.Println("rectangle perimeter:", r1.perimeter())

	// fmt.Println("circle volume:", c1.(circle).volume())

	print("circle", c1)
	print("rectangle", r1)
}
