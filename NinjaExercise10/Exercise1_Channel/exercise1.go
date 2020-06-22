package main

import (
	"fmt"
)

func main() {
	//Buffer = ,1)
	c := make(chan int, 1)
	c <- 42

	/* func literal
	go func() {
		c <- 42
	}()
	*/

	fmt.Println(<-c)
}
