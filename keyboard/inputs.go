package keyboard

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var number1 int
var number2 int
var leyend string
var err error

func inputNumber(number string, scanner bufio.Scanner, variable *int) {
	fmt.Println("Ingresa el número " + number + ": ")

	if scanner.Scan() {
		*variable, err = strconv.Atoi(scanner.Text())

		if err != nil {
			panic("El dato ingresado no es un número:" + err.Error())
		}
	}
}

func InputNumbers() {
	scanner := bufio.NewScanner(os.Stdin)

	inputNumber("1", *scanner, &number1)
	inputNumber("2", *scanner, &number2)

	fmt.Println("Ingresa el leyenda: ")
	if scanner.Scan() {
		leyend = scanner.Text()
	}

	fmt.Println(leyend, number1*number2)
}
