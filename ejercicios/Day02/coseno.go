/*
Se lee desde el teclado el valor de un angulo en grados. Se desea mostrar el coseno y el seno.
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	var angulo, coseno, seno float64
	fmt.Println("Ingresa el angulo en grados: ")
	fmt.Scan(&angulo)
	coseno = math.Cos(angulo * math.Pi / 180)
	seno = math.Sin(angulo * math.Pi / 180)
	fmt.Printf("El coseno de %.2f grados es %.2f\n", angulo, coseno)
	fmt.Printf("El seno de %.2f grados es %.2f\n", angulo, seno)
}
