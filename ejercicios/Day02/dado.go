/*
Mostrar en letras el número de la cara de un dado obtenido al azar.
*/
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	dado := rand.Intn(6) + 1 // Genera un número aleatorio entre 1 y 6
	switch dado {
	case 1:
		fmt.Println("Uno")
	case 2:
		fmt.Println("Dos")
	case 3:
		fmt.Println("Tres")
	case 4:
		fmt.Println("Cuatro")
	case 5:
		fmt.Println("Cinco")
	case 6:
		fmt.Println("Seis")
	}
}
