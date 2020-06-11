package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {
	x := person{
		first: "James",
		last:  "Bond",
		age:   42,
	}

	fmt.Printf("%T\n", &x)

	changeMe(&x)

}

func changeMe(p *person) {
	fmt.Println("Original:\t", p.first)
	p.first = "name changed"
	fmt.Println("Modified:\t", p.first)
}
