package functions

import "fmt"

func Calculations() {
	var number3 int = 32
	var number4 int = 243

	calculate := func(number1 int, number2 int) int {
		return number1 + number2 + number3 + number4
	}

	fmt.Printf("%d\n", calculate(10, 25))

	calculate = func(number1 int, number2 int) int {
		return number1 * number2 * number3
	}

	fmt.Printf("%d\n", calculate(10, 25))
}
