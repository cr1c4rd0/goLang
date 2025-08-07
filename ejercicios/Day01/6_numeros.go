/*
Ingresar dos números por teclado y sumarlos. En caso que los números sean negativos, previo a la suma se debe cambiar su signo
*/
package main

import "fmt"

func main() {
	var num1, num2 int
	fmt.Println("Ingrese el primer número: ")
	fmt.Scan(&num1)
	fmt.Println("Ingrese el segundo número: ")
	fmt.Scan(&num2)
	if num1 < 0 {
		num1 = -num1
		fmt.Println("El primer número era negativo, se ha cambiado su signo.")
	}
	if num2 < 0 {
		num2 = -num2
		fmt.Println("El segundo número era negativo, se ha cambiado su signo.")
	}
	suma := num1 + num2
	fmt.Println("La suma es:", suma)
}
