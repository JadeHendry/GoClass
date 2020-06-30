package main

import (
	dog "GoClass/NinjaExercise12/dog_Documentation"
	"fmt"
)

type canine struct {
	name string
	age  int
}

func main() {
	humanYears := 7
	fmt.Printf("My dog is %v years old.\n", humanYears)

	dogYears := dog.Years(humanYears)
	fido := canine{"fido", dogYears}
	fmt.Printf("That's %v in dog years.\n", dogYears)
	fmt.Println(fido)
}
