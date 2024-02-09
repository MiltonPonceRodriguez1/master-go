package defer_panic_recovery

import (
	"fmt"
	"log"
)

func ShowDefer() {
	fmt.Println("Primer mensaje")

	defer fmt.Println("Mensaje Final")

	fmt.Println("Segundo mensaje")
}

func ShowPanic() {
	a := 1

	defer func() {
		if reco := recover(); reco != nil {
			log.Fatalf("Error: %v", reco)
		}
	}()

	if a == 1 {
		panic("Se encontro el Valor 1")
	}
}
