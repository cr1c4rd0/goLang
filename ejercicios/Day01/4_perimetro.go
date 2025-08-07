/*
Mostrar el perímetro de una circuferencia, siempre y cuando el radio que se ingresa sea mayor a cero (controlar dicho ingreso)
*/
package main

import "fmt"

func main() {
	var radio float64
	fmt.Println("ingresa el valor R del radio: ")
	fmt.Scan(&radio)
	if radio > 0 {
		perimetro := 2 * 3.14 * radio
		fmt.Println("El perímetro de la circunferencia es:", perimetro)
	} else {
		fmt.Println("El radio debe ser mayor a cero.")
	}
}
