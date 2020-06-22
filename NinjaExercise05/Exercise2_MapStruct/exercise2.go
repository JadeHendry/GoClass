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

	m := map[string]person{
		jh.first: jh,
		sh.first: sh,
	}

	fmt.Println(jh)
	fmt.Println(sh)
	fmt.Println(m)

	fmt.Printf("%v %v:\n", jh.first, jh.last)
	for i, v := range jh.iceCream {
		fmt.Println("\t", i, v)
	}

	fmt.Printf("%v %v:\n", sh.first, sh.last)
	for i, v := range sh.iceCream {
		fmt.Println("\t", i, v)
	}

	for k, v := range m {
		fmt.Printf("%v\n", k)
		fmt.Println("\t first:", v.first, "Last:", v.last)
		for i, v2 := range v.iceCream {
			fmt.Println("\t", i, v2)
		}
	}
}
