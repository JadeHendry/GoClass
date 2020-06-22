package main

import "fmt"

func main() {
	defer foo()
	fmt.Println("Hello World!")
}

func foo() {
	//anonymous function
	defer func() {
		fmt.Println("Defer worked if after foo ran")
	}()
	fmt.Println("Foo ran")
}
