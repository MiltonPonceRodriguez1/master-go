package arrays_slices

import "fmt"

var tableS []int = []int{2, 5, 4}
var array [10]int = [10]int{6, 98, 21, 674, 18, 36, 78, 9}

func ShowSlice() {
	fmt.Println(tableS)

	part1 := array[3:]  // Slice creado desde un vector, desde la posicion 3 [3, n-1]
	part2 := array[:5]  // Slice creado desde un vector, desde la posicion 0 hasta la 5 [0, 5)
	part3 := array[6:7] // Slice creado desde un vector, desde la posicion 6 hasta la 7 [6, 7)

	fmt.Println(part1)
	fmt.Println(part2)
	fmt.Println(part3)

}

func Capacity() {
	elements := make([]int, 5, 20)
	fmt.Printf("Longitud %d, Capacidad %d\n", len(elements), cap(elements))

	nums := make([]int, 0, 0)

	for i := 0; i < 100; i++ {
		nums = append(nums, i)
	}

	fmt.Printf("Longitud %d, Capacidad %d\n", len(nums), cap(nums))

}
