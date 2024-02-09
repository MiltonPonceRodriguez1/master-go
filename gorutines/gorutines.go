package gorutines

import (
	"fmt"
	"strings"
	"time"
)

func MySlowName(name string) {
	letras := strings.Split(name, "")

	for _, letra := range letras {
		time.Sleep(1 * time.Second)
		fmt.Println(letra)
	}
}
