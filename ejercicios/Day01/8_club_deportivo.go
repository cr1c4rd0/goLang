/*
Se ingresa por teclado la categoría de un socio del club deportivo sol naciente y su antiguedad en años.
Las categorías posibles son A, B y C.
Luego se desea saber si el socio ingresado tiene Categoría A o su antiguedad se encuentra emtre los 10 y 20 años,
en esos casos se pide mostrar un mensaje qie exprese lo siguiente: "Socio VIP".
*/
package main

import "fmt"

func main() {
	var categoria string
	var antiguedad int
	fmt.Println("Ingrese su categoría: ")
	fmt.Scan(&categoria)
	fmt.Println("Ingrese su antigüedad en años: ")
	fmt.Scan(&antiguedad)
	if categoria == "a" || antiguedad >= 10 && antiguedad <= 20 {
		fmt.Println("Eres Socio VIP")
	} else {
		fmt.Println("No eres Socio VIP")
	}
}
