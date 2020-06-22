package main

import "fmt"

func main() {
	x := 0
	fmt.Printf("Type: %T\nValue: %v\n", x, x)
	fmt.Printf("Type: %T\nValue: %v\n", &x, &x)
}
