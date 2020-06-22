package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println(runtime.GOARCH)
	fmt.Println("Begin CPU\t", runtime.NumCPU())
	fmt.Println("Begin GoR\t", runtime.NumGoroutine())

	wg.Add(3)

	go bar()
	fmt.Println("Mid CPU\t", runtime.NumCPU())
	fmt.Println("Mid GoR\t", runtime.NumGoroutine())
	go foo()
	fmt.Println("Mid CPU\t", runtime.NumCPU())
	fmt.Println("Mid GoR\t", runtime.NumGoroutine())
	go foobar()
	fmt.Println("Mid CPU\t", runtime.NumCPU())
	fmt.Println("Mid GoR\t", runtime.NumGoroutine())

	wg.Wait()
	fmt.Println("End CPU\t", runtime.NumCPU())
	fmt.Println("End GoR\t", runtime.NumGoroutine())

}

func bar() {
	fmt.Println("Goroutine1")
	wg.Done()
}

func foo() {
	fmt.Println("Goroutine2")
	wg.Done()
}

func foobar() {
	fmt.Println("Goroutine3")
	wg.Done()
}
