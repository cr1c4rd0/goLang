package main

import "fmt"

func dividir(dividendo, divisor int) int {
	validateCero(divisor)
	fmt.Println("resultado: ", dividendo/divisor)
	return dividendo / divisor

}

func validateCero(divisor int) {
	if divisor == 0 {
		panic("no se puede dividir por cero")
	}
}

func main() {
	dividir(100, 0)

}
