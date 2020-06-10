package main

import "fmt"

func main() {
	bd := 1995
	i := 0
	for {
		if bd <= 2020 {
			fmt.Println(bd)
			bd++
			i++
		} else {
			break
		}
	}

}
