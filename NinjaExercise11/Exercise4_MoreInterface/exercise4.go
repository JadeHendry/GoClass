package main

import (
	"fmt"
	"log"
)

type sqrtError struct {
	lat  string
	long string
	err  error
}

func (se sqrtError) Error() string {
	return fmt.Sprintf("math error: %v %v %v", se.lat, se.long, se.err)
}

func main() {

	_, err := sqrt(-10.23)
	if err != nil {
		log.Println(err)
	}

}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		//this allows me to pass the value to show what was wrong
		e := fmt.Errorf("Cannot sqare %v because it is negative", f)
		err := sqrtError{"50.2289 N", "99.4656 W", e}
		/*
			err := sqrtError{
				lat:  "50.2289 N",
				long: "99.4656 W",
				err:  errors.New("- cannot sqaure a negative -"),
			}
		*/
		return 0, err
	}
	return 42, nil
}
