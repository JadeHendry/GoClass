// Package dog allows us to more fully understand dogs.
package dog

// Years converts human years into dog years.
func Years(i int) int {
	return i * 7
}

type canine struct {
	name string
	age  int
}
