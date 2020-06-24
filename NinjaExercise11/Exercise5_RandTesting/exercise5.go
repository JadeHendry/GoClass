package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	for i := 0; i < 100; i++ {
		fmt.Print(rand.Intn(100), ",")
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		fmt.Print(r1.Intn(100), "\t")

		fmt.Print(rand.Intn(100), ",")
		fmt.Println(r1.Intn(100))
	}
}
