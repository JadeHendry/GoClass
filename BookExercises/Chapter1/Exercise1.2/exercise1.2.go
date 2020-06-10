package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	test := func() {
		for i, v := range os.Args[:] {
			fmt.Println(i, v)
		}
	}
	fmt.Println(timeToExecute(test))
}

func timeToExecute(f func()) float64 {
	start := time.Now()
	f()
	return time.Since(start)
}
