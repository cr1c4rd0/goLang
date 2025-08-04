package main

import (
	"fmt"
	"math"
)

func main() {
	var base float64 = 0
	var altura float64 = 0

	fmt.Println("Ingrese la base del triángulo:")
	fmt.Scanln(&base)
	fmt.Println("Ingrese la altura del triángulo:")
	fmt.Scanln(&altura)

	var area = (base * altura) / 2
	fmt.Println("El área del triángulo es:", area)

	hipotenusa := math.Sqrt(math.Pow(base, 2) + math.Pow(altura, 2))
	fmt.Println("La hipotenusa del triángulo es:", hipotenusa)

}
