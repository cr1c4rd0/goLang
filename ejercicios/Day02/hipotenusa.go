/*
Se ingresa por teclado los catetos de un triángulo rectángulo. Se desea hallar y mostrar su hipotenusa
*/
package main

import "fmt"

func main() {
	var cateto1, catateto2 float64
	fmt.Println("Ingrese el cateto A:")
	fmt.Scan(&cateto1)
	fmt.Println("Ingrese el cateto B:")
	fmt.Scan(&catateto2)
	pitagoras := (cateto1 * cateto1) + (catateto2 * catateto2)
	hipotenusa := pitagoras * 0.5
	fmt.Printf("La hipotenusa es: %.2f\n", hipotenusa)
}
