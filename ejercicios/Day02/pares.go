/*
Ingresa un número natural por teclado. Se desea saber y mostrar si es par o impar.
*/
package main

import "fmt"

func main() {
	var num int
	fmt.Println("Ingrea un número natural: ")
	fmt.Scan(&num)
	if num%2 == 0 {
		fmt.Println("El número es par")
	} else {
		fmt.Println("El número es impar")
	}
}
