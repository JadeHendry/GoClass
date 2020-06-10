package main

import "fmt"

func main() {
	f := funcReturn()
	f()
}

func funcReturn() func() {
	return func() {
		fmt.Println("This is the returned func")
	}
}
