package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("archivo.txt")
	if err != nil {
		fmt.Println("Error al crear el archivo:", err)
		return
	}

	defer file.Close()

	_, err = file.WriteString("¡Hola, mundo desde un archivo!\n")
	if err != nil {
		fmt.Println("Error al escribir en el archivo:", err)
		return
	}

}
