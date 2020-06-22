package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type square struct {
	len float64
}

type shape interface {
	area() float64
}

func main() {
	circ := circle{
		radius: 5,
	}

	squa := square{
		len: 5,
	}

	info(circ)
	info(squa)
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (s square) area() float64 {
	return math.Pow(s.len, 2)
}

func info(s shape) {
	fmt.Println(s.area())
}
