/*
Ingresar un nombre y un apelido en distintas variables y luego mostrar en forma concatenada
el nombre seguido del apellido. Por ejemplo:
Si el nombre es "Ana" y el apellido es "Perez", la salida sería "Ana Perez"
*/
package main

import "fmt"

func main() {
	var nombre, apellido string
	fmt.Println("Ingrese su nombre: ")
	fmt.Scan(&nombre)
	fmt.Println("Ingrese su apellido: ")
	fmt.Scan(&apellido)
	fmt.Println("Bienvenido " + nombre + " " + apellido)
	return
}
