package main

import "master-go/defer_panic_recovery"

// "fmt"
// "master-go/exercises"
// "master-go/keyboard"
// "runtime"
// "master-go/iterations"
// "master-go/files"
// "master-go/arrays_slices"
// "master-go/maps"
// e "master-go/interfaces_exercises"
// "master-go/models"
// "master-go/users"
// "master-go/functions"

func main() {
	// state, text := variables.ConvertToText(1588)
	// fmt.Println(state, text)

	// if os := runtime.GOOS; os == "linux." || os == "OS X." {
	// 	fmt.Println("Esto no es Windows, es ", os)
	// } else {
	// 	fmt.Println("Esto es Windows")
	// }

	// switch os := runtime.GOOS; os {
	// case "linux":
	// 	fmt.Println("Esto es Linux")
	// case "darwin":
	// 	fmt.Println("Esto es Darwin")
	// default:
	// 	fmt.Printf("%s \n", os)
	// }

	// intSend, textMsg := exercises.ConvertToInt("104x")

	// fmt.Println(intSend, textMsg)

	// keyboard.InputNumbers()

	// iterations.Loop()

	// fmt.Println(exercises.PrintMultiplicationTable())

	// ? Manejo de Archivos
	// files.WriteTable()
	// files.ConcatTable()

	// files.ReadFile()
	// files.ReadFileTwo()

	// ? Recursion y Clousures
	// functions.Calculations()
	// functions.CallClosure()
	// functions.Exponent(1)

	// ? Slices
	// arrays_slices.ShowArrays()
	// arrays_slices.Capacity()

	// ? Mapas (Arreglos asociativos)
	// maps.ShowMaps()

	// ? Modelos
	// users.UserRegister()

	// ? Herencia, Polimorfismo e Interfaces
	// milton := new(models.Man)
	// e.HumansBreathing(milton)

	// ana := new(models.Women)
	// e.LivingBeingBreathing(ana)

	// ? Defer, Panic y Recovery
	defer_panic_recovery.ShowPanic()

}
