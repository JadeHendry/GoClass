package main

import "fmt"

type person struct {
	first    string
	last     string
	iceCream []string
}

func main() {
	jh := person{
		first: "Jade",
		last:  "Hendry",
		iceCream: []string{
			"Mint Chip",
			"Cookie Dough",
			"Smores",
		},
	}

	sh := person{
		first: "Savannah",
		last:  "Hendry",
		iceCream: []string{
			"Double Chocolate",
			"Split",
			"Chocolate Raspberry",
		},
	}

	fmt.Println(jh)
	fmt.Println(sh)

	fmt.Printf("%v %v:\n", jh.first, jh.last)
	for i, v := range jh.iceCream {
		fmt.Println("\t", i, v)
	}

	fmt.Printf("%v %v:\n", sh.first, sh.last)
	for i2, v2 := range sh.iceCream {
		fmt.Println("\t", i2, v2)
	}
}
