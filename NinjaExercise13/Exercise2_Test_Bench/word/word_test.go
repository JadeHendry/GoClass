package word

import (
	"GoClass/NinjaExercise13/Exercise2_Test_Bench/quote"
	"fmt"
	"reflect"
	"testing"
)

func TestUseCount(t *testing.T) {
	m := map[string]int{
		"One":   1,
		"Two":   2,
		"Three": 3,
		"Four":  4,
	}
	s1 := "One Two Two Three Three Three Four Four Four Four"
	m1 := UseCount(s1)
	if !reflect.DeepEqual(m, m1) {
		t.Fatal("Map lengths differ.")
	}
}

func TestCount(t *testing.T) {
	s := "Mama mia. Here we go again."
	c := Count(s)
	if c != 6 {
		t.Fatal("Did not count words correctly")
	}
}

func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseCount(quote.SunAlso)
	}
}

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count(quote.SunAlso)
	}
}

func ExampleUseCount() {
	s1 := "One Two Two Three Three Three Four Four Four Four"
	m1 := UseCount(s1)
	fmt.Println(m1)
	//Output:
	//map[Four:4 One:1 Three:3 Two:2]
}

func ExampleCount() {
	fmt.Println(Count("Mama mia. Here we go again."))
	//Output:
	//6
}

//for coverage use command "go test cover"

//then to display the coverage in a webpage
//"go test -coverprofile c.out"
//"go tool cover -html=c.out"
