package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var mtex sync.Mutex

func main() {
	incrementor := 0
	gs := 100
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			mtex.Lock()
			v := incrementor
			v++
			incrementor = v
			fmt.Println(incrementor)
			mtex.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("end value:", incrementor)
}
