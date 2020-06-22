package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {
	jb := person{
		first: "James",
		last:  "Bond",
		age:   42,
	}

	jb.speak()
}

func (p person) speak() {
	fmt.Println("Hi, my name is", p.first, p.last, "and I am", p.age, "years old.")
}
