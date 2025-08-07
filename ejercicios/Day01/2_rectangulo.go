/*
Hallar el perímetro y el área de un rectángulo ingresando la base b y la altura h
*/
package main

import "fmt"

func main() {
	var base, altura float64
	fmt.Println("Ingrese  la base B del rectangulo: ")
	fmt.Scan(&base)
	fmt.Println("Ingrese la altura H del rectangulo: ")
	fmt.Scan(&altura)
	perimetro := 2 * (base + altura)
	area := base * altura
	fmt.Printf("El perimetro del rectangulo es: %.2f\n", perimetro)
	fmt.Printf("El area del rectangulo es: %.2f\n", area)
}
