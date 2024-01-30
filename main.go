package main

import (
	"fmt"
	"master-go/exercises"
	"runtime"
)

func main() {
	// state, text := variables.ConvertToText(1588)
	// fmt.Println(state, text)

	if os := runtime.GOOS; os == "linux." || os == "OS X." {
		fmt.Println("Esto no es Windows, es ", os)
	} else {
		fmt.Println("Esto es Windows")
	}

	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("Esto es Linux")
	case "darwin":
		fmt.Println("Esto es Darwin")
	default:
		fmt.Printf("%s \n", os)
	}

	intSend, textMsg := exercises.ConvertToInt("1kj50")

	fmt.Println(intSend, textMsg)
}
