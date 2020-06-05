package main

import "fmt"

func main() {
	mp := []string{"Miss", "Moneypenny", "Helloooooo, James"}
	jb := []string{"James", "Bond", "Shaken, not stirred"}

	x := [][]string{jb, mp}

	fmt.Println(x)

	for i, slice := range x {
		fmt.Println("record:", i)
		for index, v := range slice {
			fmt.Println("\tindex position:", index, "\tvalue:", v)
		}
	}

}
