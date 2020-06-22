package main

import "fmt"

func main() {
	favSport := "snowboarding"

	switch favSport {
	case "skiing":
		fmt.Println("My favorite sport is skiing")
	case "surfing":
		fmt.Println("My favorite sport is surfing")
	case "snowboarding":
		fmt.Println("My favorite sport is snowboarding")
	case "sky diving":
		fmt.Println("My favorite sport is sky diving")
	}
}
