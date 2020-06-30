package mymath

import (
	"fmt"
	"testing"
)

func TestCenteredAvg(t *testing.T) {

	type test struct {
		data   []int
		answer float64
	}

	xtest := []test{
		test{[]int{10, 20, 40, 60, 80}, 40},
		test{[]int{2, 4, 6, 8, 12}, 6},
		test{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 5},
	}

	for _, v := range xtest {
		f := CenteredAvg(v.data)
		if f != v.answer {
			t.Fatalf("Centered average is not correct: %v != %v", v, v.answer)
		}
	}

	/* original idea. Above is better
	a := []int{1, 4, 6, 8, 100}
	if CenteredAvg(a) != 6.0 {
		t.Fatal("func CenteredAvg does not work properly")
	}
	b := []int{0, 8, 10, 1000}
	if CenteredAvg(b) != 9.0 {
		t.Fatal("func CenteredAvg does not work properly")
	}
	c := []int{9000, 4, 10, 8, 6, 12}
	if CenteredAvg(c) != 9.0 {
		t.Fatal("func CenteredAvg does not work properly")
	}
	*/
}

func BenchmarkCenteredAvg(b *testing.B) {
	a := []int{1, 4, 6, 8, 100}
	for i := 0; i < b.N; i++ {
		CenteredAvg(a)
	}
}

func ExampleCenteredAvg() {
	a := []int{1, 4, 6, 8, 100}
	fmt.Println(CenteredAvg(a))
	//Output:
	//6
}
