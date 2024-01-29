package main

import (
	"fmt"
	"master-go/variables"
)

func main() {
	state, text := variables.ConvertToText(1588)

	fmt.Println(state, text)
}
