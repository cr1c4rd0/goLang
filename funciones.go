package main

import (
	"fmt"
)

func main() {
	saludo := hello("Cristian")
	fmt.Println(saludo)
	suma, mul := calc(3, 4)
	fmt.Println("La suma es: ", suma)
	fmt.Println("La multiplicación es: ", mul)
}

func hello(name string) string {
	return "Hello," + name
}

func calc(a, b int) (sum, mul int) {
	sum = a + b
	mul = a * b
	return sum, mul
}
