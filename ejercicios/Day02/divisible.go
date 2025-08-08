/*
Ingresar un número entero para saber si es divisible por 7 y es mayor a 40.
*/
package main

import "fmt"

func main() {
	var number int
	fmt.Println("Ingresa un número entero: ")
	fmt.Scan(&number)
	if number%7 == 0 && number > 40 {
		fmt.Println("El número es divisible por 7 y es mayor a 40.")
	} else {
		fmt.Println("El número no es divisible por 7 o no es mayor a 40.")
	}
}
