package functions

import "fmt"

func Exponent(value int) {
	if value > 1000 {
		return
	}

	fmt.Printf("%d \n", value)
	Exponent(value * 2)
}
