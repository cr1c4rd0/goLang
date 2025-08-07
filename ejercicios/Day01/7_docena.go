/*
Solicitar al usuario un número natural y verificar que el número ingresado se encuentre dentro de la primera docena de los número naturales,
es decir entre el 1 y el 12.
*/
package main

import "fmt"

func main() {
	var num int
	fmt.Println("Ingrese un número natural: ")
	fmt.Scan(&num)
	if num < 1 || num > 12 {
		fmt.Println("El número ingresado no es valido, ya que no está entre 1 y 12 en la primera docena de los números naturales.")
	} else {
		fmt.Println(num, ": es un número natural válido.")
	}
}
