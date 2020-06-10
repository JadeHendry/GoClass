package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	fmt.Println(len(os.Args))

	test := func() {
		for i, v := range os.Args[:] {
			fmt.Println(i, v)
		}
	}
	fmt.Println(timeToExecute(test))

}

func timeToExecute(f func()) time.Duration {
	start := time.Now()
	f()
	return time.Since(start)
}
