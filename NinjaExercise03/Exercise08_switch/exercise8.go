package main

import "fmt"

func main() {
	x := 15
	switch {
	case x < 15:
		fmt.Println("x is less than 15")
	case x == 15:
		fmt.Println("x equals 15")
	case x > 15:
		fmt.Println("x is greater than 15")
	}
}
