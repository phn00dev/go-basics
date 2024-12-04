package main

import (
	"fmt"
	"math"
)

type Circle struct {
	radius float64
}

type Rectangle struct {
	width, height float64
}

func (c Circle) Area() float64 {

	return c.radius * c.radius * math.Pi
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func main() {
	c1 := Circle{10}
	c2 := Circle{25}
	r1 := Rectangle{9, 4}
	r2 := Rectangle{12, 2}

	fmt.Println("Area of c1 is: ", c1.Area())
	fmt.Println("Area of c2 is: ", c2.Area())
	fmt.Println("Area of r1 is: ", r1.Area())
	fmt.Println("Area of r2 is: ", r2.Area())
}
