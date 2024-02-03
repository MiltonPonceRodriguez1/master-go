package functions

import "fmt"

func table(value int) func() int {
	number := value
	sequence := 0

	return func() int {
		sequence++
		return number * sequence
	}
}

func CallClosure() {
	numberTable := 2
	MyTable := table(numberTable)

	for i := 1; i <= 10; i++ {
		fmt.Printf("%d\n", MyTable())
	}
}
