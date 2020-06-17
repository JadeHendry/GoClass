package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type human interface {
	speak()
}

func main() {
	jade := person{
		first: "Jade",
		last:  "Hendry",
		age:   24,
	}

	jade.speak()
	saySomething(&jade)
}

func (p *person) speak() {
	fmt.Println(p.first, p.last, "- Age:", p.age)
}

func saySomething(h human) {
	h.speak()
}
