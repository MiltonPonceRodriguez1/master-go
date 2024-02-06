package arrays_slices

import "fmt"

var table [10]int
var matrix [20][30]int

func ShowArrays() {
	table[7] = 33
	table[2] = 54

	table2 := [10]string{"Pablo", "Juan"}

	fmt.Println(table)
	fmt.Println(table2)

	for i := 0; i < len(table); i++ {
		fmt.Printf("[%d] ", table[i])
	}

	matrix[7][24] = 15
	matrix[1][2] = 5
	fmt.Println()

	for i := 0; i < 20; i++ {
		for j := 0; j < 30; j++ {
			fmt.Printf("[%d] ", matrix[i][j])
		}
		fmt.Println()
	}
}
