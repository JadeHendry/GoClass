package main

import (
	"GoClass/NinjaExercise13/Exercise2_Test_Bench/quote"
	"GoClass/NinjaExercise13/Exercise2_Test_Bench/word"
	"fmt"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}
