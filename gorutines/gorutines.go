package gorutines

import (
	"fmt"
	"strings"
	"time"
)

func MySlowName(name string, chanel1 chan bool) {
	letras := strings.Split(name, "")

	for _, letra := range letras {
		time.Sleep(1 * time.Second)
		fmt.Println(letra)
	}

	chanel1 <- true
}
