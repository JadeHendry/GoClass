package dog

import (
	"fmt"
	"testing"
)

func TestYears(t *testing.T) {
	result := Years(4)
	if result != 28 {
		t.Fatal("4 * 7 did not equal 28")
	}
}

func TestYearsTwo(t *testing.T) {
	result := YearsTwo(4)
	if result != 28 {
		t.Fatal("4 * 7 did not equal 28")
	}
}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(10)
	}
}

func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(10)
	}
}

func ExampleYears() {
	fmt.Println(Years(10))
	//Output:
	//70
}

func ExampleYearsTwo() {
	fmt.Println(Years(10))
	//Output:
	//70
}
