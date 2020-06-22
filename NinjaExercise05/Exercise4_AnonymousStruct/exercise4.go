package main

import "fmt"

func main() {

	s := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first: "James",
		friends: map[string]int{
			"Miss Moneypenny": 555,
			"Dr. No":          777,
			"Jade Hendry":     999,
		},
		favDrinks: []string{
			"jak and coke",
			"rum and coke",
			"tequilla and coke",
			"Martini and coke",
		},
	}

	fmt.Println(s)
	fmt.Println(s.first, "\n\tFriends:")
	for k, v := range s.friends {
		fmt.Println("\t\t", k, "-", v)
	}
	fmt.Println("\tFavorite Drinks:")
	for i, v := range s.favDrinks {
		fmt.Println("\t\t", i, v)
	}

}
