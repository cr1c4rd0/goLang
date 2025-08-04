/*
Supongamos que se proporciona una secuencia de número, tales como: 5 3 0 2 4 4 0 0 2 3 6 0 2
y se desea contar e imprimir el número de ceros de la secuencia.
*/
package main

import (
	"fmt"
)

func main() {
	// Supongamos que se proporciona una secuencia de números
	sec := [...]int{5, 3, 0, 2, 4, 4, 0, 0, 2, 3, 6, 0, 2}
	cantidadCeros := 0

	for _, num := range sec {
		if num == 0 {
			cantidadCeros++
		}
	}

	fmt.Println("La cantidad de ceros en la secuencia es:", cantidadCeros)
}
