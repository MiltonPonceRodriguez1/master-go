package exercises

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func inputNumber(scanner bufio.Scanner) int {
	for {
		fmt.Print("Ingresa el número de la tabla a multiplicar: ")

		if scanner.Scan() {
			if num, err := strconv.Atoi(scanner.Text()); err == nil {
				return num
			}
		}
	}

}

func PrintMultiplicationTable() string {
	scanner := bufio.NewScanner(os.Stdin)
	number := inputNumber(*scanner)

	var text string

	for i := 1; i <= 10; i++ {
		text += fmt.Sprintf("%d x %d = %d\n", number, i, number*i)
	}

	return text
}
