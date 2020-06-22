package main

import "fmt"

func main() {
	x, y := bar()
	y += foo()
	fmt.Println(x, y)

}

func foo() int {
	return 42
}

func bar() (string, int) {
	return "You foo'd this bar", 6
}
