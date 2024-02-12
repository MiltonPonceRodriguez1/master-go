package middleware

import "fmt"

func add(a, b int) int {
	return a + b
}

func subtract(a, b int) int {
	return a - b
}

func product(a, b int) int {
	return a * b
}

func MyMiddleware() {
	fmt.Println("Inicio")

	result := middlewareOperations(add)(2, 3)
	println(result)

	result = middlewareOperations(subtract)(2, 3)
	println(result)

	result = middlewareOperations(product)(2, 3)
	println(result)
}

func middlewareOperations(f func(int, int) int) func(int, int) int {
	return func(a, b int) int {
		fmt.Println("Inicio de opreación")
		return f(a, b)
	}
}
