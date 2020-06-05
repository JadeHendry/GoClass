package main

import "fmt"

type vehicle struct {
	doors string
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {

	f150 := truck{
		vehicle: vehicle{
			doors: "two",
			color: "red",
		},
		fourWheel: true,
	}

	carolla := sedan{
		vehicle: vehicle{
			doors: "four",
			color: "white",
		},
		luxury: false,
	}

	fmt.Println(f150)
	fmt.Println(carolla)
	fmt.Println(f150.doors, f150.color)
	fmt.Println(carolla.doors, carolla.color)

}
