package main

import "fmt"

func main() {
	x := 42
	y := 43
	var tft bool = x == y
	var f bool = x >= y
	var t bool = x <= y
	var tt bool = x < y
	var tf bool = x > y

	fmt.Println(tft)
	fmt.Println(t)
	fmt.Println(f)
	fmt.Println(tt)
	fmt.Println(tf)

}
