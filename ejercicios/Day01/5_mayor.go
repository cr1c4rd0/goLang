/*
Mostrar el número más grande (entre dos) ingresado por teclado.
Si los dos número son iguales mostrarl el mensaje "Son iguales"
*/
package main

import "fmt"

func main() {
	var num1, num2 int
	fmt.Println("Ingrse el primer número: ")
	fmt.Scan(&num1)
	fmt.Println("Ingrse el segundo número: ")
	fmt.Scan(&num2)
	if num1 > num2 {
		fmt.Println("El número mayor es:", num1)
	} else if num2 > num1 {
		fmt.Println("El número mayor es:", num2)
	} else {
		fmt.Println("Son iguales")
	}
}
