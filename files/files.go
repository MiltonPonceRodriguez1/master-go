package files

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"master-go/exercises"
	"os"
)

var filename string = "./files/txts/table.txt"

func WriteTable() {
	text := exercises.PrintMultiplicationTable()
	file, err := os.Create(filename)
	if err != nil {
		fmt.Printf("Error al crear el archivo: %s", err.Error())
		return
	}

	fmt.Fprintln(file, text)
	file.Close()
}

func ConcatTable() {
	text := exercises.PrintMultiplicationTable()

	if !Append(filename, text) {
		fmt.Printf("Error al concatenar el contenido !\n")
	}
}

func Append(file string, text string) bool {
	fileToOpen, err := os.OpenFile(file, os.O_WRONLY|os.O_APPEND, 0644)

	if err != nil {
		fmt.Printf("Error al abrir el fichero: %s!\n", err.Error())
		return false
	}

	if _, err = fileToOpen.WriteString(text + "\n"); err != nil {
		fmt.Printf("Error al escribir el fichero: %s!\n", err.Error())
		return false
	}

	fileToOpen.Close()
	return true
}

func ReadFile() {
	file, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Printf("Error al leer el archivo: %s", err.Error())
	}

	fmt.Printf("%s", string(file))
}

func ReadFileTwo() {
	file, err := os.Open(filename)

	if err != nil {
		fmt.Printf("Error al leer el archivo: %s", err.Error())
		return
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		register := scanner.Text()
		fmt.Printf("> %s \n", register)
	}

	file.Close()
}
