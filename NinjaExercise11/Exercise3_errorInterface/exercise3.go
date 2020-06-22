package main

import "fmt"

type customErr struct {
	str string
}

func main() {
	c1 := customErr{
		str: "Need more coffee.",
	}

	foo(c1)
}

func foo(e error) {
	fmt.Println("foo --", e)
}

func (ce customErr) Error() string {
	return fmt.Sprintf("Here is the error: %v", ce.str)
}
