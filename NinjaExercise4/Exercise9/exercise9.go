package main

import "fmt"

func main() {
	x := map[string][]string{
		"bond_james":      []string{"Shaken, not stirred", "Martinis", "Women"},
		"moneypenny_miss": []string{"James Bond", "Literature", "Computer Science"},
		"no_dr":           []string{"Being evil", "Ice cream", "Sunsets"},
	}

	x["flemming_ian"] = []string{"steaks", "cigarrs", "espionage"}

	for k, v := range x {
		fmt.Println("This is the record for", k)
		for i, val := range v {
			fmt.Printf("\t%v %v\n", i, val)
		}
	}
}
